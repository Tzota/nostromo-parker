
#!/usr/bin/env bash
set -euo pipefail

#######################
ARCH=amd64
 while :; do
    case "${1-}" in
    -v | --verbose) set -x ;;
    -a | --arch) # example named parameter
      ARCH="${2-}"
      shift
      ;;
    -?*) die "Unknown option: $1" ;;
    *) break ;;
    esac
    shift
  done
#######################
