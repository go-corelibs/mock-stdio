#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += STDIO
STDIO_MK_SUMMARY := go-corelibs/mock-stdio
STDIO_MK_VERSION := v1.0.0

include CoreLibs.mk
