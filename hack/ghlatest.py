#!/usr/bin/env python3
# -*- coding: utf-8 -*-
#  Copyright (c) 2019-2021 Prasad Tengse
#

import argparse
import json
import sys
from urllib.request import urlopen


def main(user, repo):
    url = f"https://api.github.com/repos/{user}/{repo}/releases/latest"
    with urlopen(url=url) as response:
        if response.status is not 200:
            sys.exit(1)
        r = json.load(response)
        print(f"{r['tag_name']}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Get tagname of latest GitHub release for a repository",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
        add_help=True,
    )
    parser.add_argument(
        "-r", "--repo", required=True, type=str, help="GitHub Repository"
    )
    parser.add_argument(
        "-user", "--user", required=True, type=str, help="GitHub Username"
    )

    args = parser.parse_args()
    main(repo=args.repo, user=args.user)
