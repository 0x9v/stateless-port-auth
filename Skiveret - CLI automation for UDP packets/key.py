import subprocess
import argparse
import time
import sys


KNOCK_SEQ = [7000, 9000, 8000]
BAD_SEQ = [8000, 8000, 8000]
BIN_PL = './knock'
DUMMY_PL = 'eaccent'


def launch_seq():
    try:
        for port in KNOCK_SEQ:
            print(f'[...] firing payload at {port}...')
            command = [BIN_PL, args.target, str(port), 'eaccent']
            subprocess.run(command)
            time.sleep(0.5)
    except FileNotFoundError:
        print(f'[!!!] error: binary payload {BIN_PL} not found')
        sys.exit(1)


def test_failure():
    try:
        for port in BAD_SEQ:
            print(f'[xxx] firing out-of-order payload at {port}...')
            command = [BIN_PL, args.target, str(port), 'eaccent']
            subprocess.run(command)
            time.sleep(0.5)
    except FileNotFoundError:
        print(f'[!!!] error: binary payload {BIN_PL} not found')
        sys.exit(1)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='stealth port knocking cli tool')
    parser.add_argument('-t', '--target',
                        default='127.0.0.1',
                        help='the target ip address')

    parser.add_argument('-m', '--mode',
                        required=True,
                        choices=['exploit', 'fuzz'],
                        help='modes: exploit (valid sequence) fuzz (bad sequence)')

    args = parser.parse_args()
    if args.mode == 'exploit':
        print(f'[...] initiating exploit mode against {args.target}')
        launch_seq()
    elif args.mode == 'fuzz':
        print(f'[???] initiating fuzz mode against {args.target}')
        test_failure()
