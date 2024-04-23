// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT
//
// This file is automatically generated by go run generate.go DO NOT EDIT.

//go:build darwin

#include "textflag.h"

GLOBL	·cf_trampoline_release_addr(SB), RODATA, $8
DATA	·cf_trampoline_release_addr(SB)/8, $cf_trampoline_release<>(SB)
TEXT    cf_trampoline_release<>(SB),NOSPLIT,$0-0
	        JMP	CoreFoundation_CFRelease(SB)
            RET

GLOBL	·cf_trampoline_array_get_count_addr(SB), RODATA, $8
DATA	·cf_trampoline_array_get_count_addr(SB)/8, $cf_trampoline_array_get_count<>(SB)
TEXT    cf_trampoline_array_get_count<>(SB),NOSPLIT,$0-0
	        JMP	CoreFoundation_CFArrayGetCount(SB)
            RET

GLOBL	·cf_trampoline_array_get_value_at_index_addr(SB), RODATA, $8
DATA	·cf_trampoline_array_get_value_at_index_addr(SB)/8, $cf_trampoline_array_get_value_at_index<>(SB)
TEXT    cf_trampoline_array_get_value_at_index<>(SB),NOSPLIT,$0-0
	        JMP	CoreFoundation_CFArrayGetValueAtIndex(SB)
            RET

GLOBL	·cf_trampoline_data_create_addr(SB), RODATA, $8
DATA	·cf_trampoline_data_create_addr(SB)/8, $cf_trampoline_data_create<>(SB)
TEXT    cf_trampoline_data_create<>(SB),NOSPLIT,$0-0
	        JMP	CoreFoundation_CFDataCreate(SB)
            RET

