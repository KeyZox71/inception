#!/bin/sh

set -e

# Define variables from environment
REPO=${BORG_REPO:-/backup}
BORG_PASSPHRASE=$(getpassphrase)
SOURCE=${BORG_SOURCE:-/source}
COMPRESSION=${BORG_COMPRESS:-zstd}
PRUNE_KEEP_DAILY=${BORG_PRUNE_KEEP_DAILY:-7}
PRUNE_KEEP_WEEKLY=${BORG_PRUNE_KEEP_WEEKLY:-4}
PRUNE_KEEP_MONTHLY=${BORG_PRUNE_KEEP_MONTHLY:-6}
EXCLUDE_PATTERNS=${BORG_EXCLUDE_PATTERNS:-}
CHECK_LAST=${BORG_CHECK_LAST}

BAK_ARGS="--compression $COMPRESSION"

if [ -z "$BORG_PASSPHRASE" ]; then
	echo "Could not found passphrase"
	exit 1
fi

if [ -n "$EXCLUDE_PATTERNS" ]; then
	BAK_ARGS="$BAK_ARGS --exclude $EXCLUDE_PATTERNS"
fi

# Borg backup command
borg create --stats $BAK_ARGS \
	$REPO::$(date +%Y-%m-%d) $SOURCE

# Borg prune command

echo "Creating backup..."
borg prune --list $REPO --keep-daily=$PRUNE_KEEP_DAILY --keep-weekly=$PRUNE_KEEP_WEEKLY --keep-monthly=$PRUNE_KEEP_MONTHLY

# Borg check command
CHECK_ARGS=""

if [ -n "$CHECK_LAST" ]; then
	CHECK_ARGS="$CHECK_ARGS --last $CHECK_LAST"
fi
if [ -n "$CHECK_DATA" ]; then
	CHECK_ARGS="$CHECK_ARGS --verify-data"
fi

borg check $CHECK_ARGS $REPO
