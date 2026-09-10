#!/usr/bin/env python3
"""Build release archives containing the node and CLI wallet only."""
import argparse
import hashlib
import os
from pathlib import Path
import platform
import re
import shutil
import subprocess
import tarfile
import zipfile

ROOT = Path(__file__).resolve().parents[1]
TARGETS = ('linux-amd64', 'windows-amd64')


def build(target, version):
    goos, goarch = target.split('-')
    name = f'nonsense-{version}-{target}'
    out = ROOT / 'dist' / name
    out.mkdir(parents=True, exist_ok=True)
    env = os.environ.copy()
    env.update(GOOS=goos, GOARCH=goarch, CGO_ENABLED='1')
    env.setdefault('GOMAXPROCS', '2')
    if goos == 'windows' and platform.system() != 'Windows':
        env['CC'] = 'x86_64-w64-mingw32-gcc'
    elif goos == 'linux' and platform.system() != 'Linux':
        raise SystemExit('Build Linux releases on Linux.')
    suffix = '.exe' if goos == 'windows' else ''
    members = []
    for executable, package in (('nonsensed', '.'), ('nonsensewallet', './cmd/nonsensewallet')):
        binary = out / (executable + suffix)
        subprocess.run([
            'go', 'build', '-mod=readonly', '-p', '2', '-trimpath', '-buildvcs=false',
            '-tags', 'netgo,osusergo', '-ldflags=-s -w -linkmode=external -extldflags=-static',
            '-o', str(binary), package,
        ], cwd=ROOT, env=env, check=True)
        members.append(binary)
    license_file = out / 'LICENSE'
    shutil.copyfile(ROOT / 'LICENSE', license_file)
    members.append(license_file)
    if goos == 'windows':
        archive = out.parent / (name + '.zip')
        with zipfile.ZipFile(archive, 'w', zipfile.ZIP_DEFLATED, compresslevel=9) as z:
            for member in members:
                z.write(member, f'{name}/{member.name}')
    else:
        archive = out.parent / (name + '.tar.gz')
        with tarfile.open(archive, 'w:gz') as t:
            for member in members:
                t.add(member, arcname=f'{name}/{member.name}')
    with archive.open('rb') as f:
        digest = hashlib.file_digest(f, 'sha256').hexdigest()
    archive.with_name(archive.name + '.sha256').write_text(f'{digest}  {archive.name}\n')
    print(f'Built {archive.name}: {digest}', flush=True)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--target', choices=TARGETS, action='append')
    parser.add_argument('--version', default='v2.3.0')
    args = parser.parse_args()
    if not re.fullmatch(r'v\d+\.\d+\.\d+(?:-[A-Za-z0-9.-]+)?', args.version):
        parser.error('version must be a release tag such as v2.3.0')
    for target in args.target or TARGETS:
        build(target, args.version)
