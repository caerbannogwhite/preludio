	.text
	.intel_syntax noprefix
	.file	"sum_float64.c"
	.globl	sum_float64_avx_intrinsics      # -- Begin function sum_float64_avx_intrinsics
	.p2align	4, 0x90
	.type	sum_float64_avx_intrinsics,@function
sum_float64_avx_intrinsics:             # @sum_float64_avx_intrinsics
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -32
	sub	rsp, 480
	mov	qword ptr [rsp + 168], rdi
	mov	qword ptr [rsp + 160], rsi
	mov	qword ptr [rsp + 152], rdx
	mov	rax, qword ptr [rsp + 160]
	mov	eax, dword ptr [rax]
	mov	edx, eax
	add	edx, 3
	test	eax, eax
	mov	ecx, eax
	cmovs	ecx, edx
	and	ecx, -4
	sub	eax, ecx
	mov	dword ptr [rsp + 148], eax
	mov	qword ptr [rsp + 176], 0
	vmovsd	xmm0, qword ptr [rsp + 176]     # xmm0 = mem[0],zero
	vmovsd	qword ptr [rsp + 456], xmm0
	vmovsd	qword ptr [rsp + 448], xmm0
	vmovsd	qword ptr [rsp + 440], xmm0
	vmovsd	qword ptr [rsp + 432], xmm0
	vmovsd	xmm1, qword ptr [rsp + 456]     # xmm1 = mem[0],zero
	vmovsd	xmm0, qword ptr [rsp + 448]     # xmm0 = mem[0],zero
	vunpcklpd	xmm1, xmm0, xmm1        # xmm1 = xmm0[0],xmm1[0]
	vmovsd	xmm2, qword ptr [rsp + 440]     # xmm2 = mem[0],zero
	vmovsd	xmm0, qword ptr [rsp + 432]     # xmm0 = mem[0],zero
	vunpcklpd	xmm2, xmm0, xmm2        # xmm2 = xmm0[0],xmm2[0]
                                        # implicit-def: $ymm0
	vmovaps	xmm0, xmm2
	vinsertf128	ymm0, ymm0, xmm1, 1
	vmovapd	ymmword ptr [rsp + 384], ymm0
	vmovapd	ymm0, ymmword ptr [rsp + 384]
	vmovapd	ymmword ptr [rsp + 96], ymm0
	mov	dword ptr [rsp + 92], 0
.LBB0_1:                                # =>This Inner Loop Header: Depth=1
	mov	eax, dword ptr [rsp + 92]
	mov	rcx, qword ptr [rsp + 160]
	mov	ecx, dword ptr [rcx]
	sub	ecx, dword ptr [rsp + 148]
	cmp	eax, ecx
	jge	.LBB0_4
# %bb.2:                                #   in Loop: Header=BB0_1 Depth=1
	mov	rax, qword ptr [rsp + 168]
	movsxd	rcx, dword ptr [rsp + 92]
	shl	rcx, 3
	add	rax, rcx
	mov	qword ptr [rsp + 184], rax
	mov	rax, qword ptr [rsp + 184]
	vmovapd	ymm0, ymmword ptr [rax]
	vmovapd	ymmword ptr [rsp + 32], ymm0
	vmovapd	ymm1, ymmword ptr [rsp + 96]
	vmovapd	ymm0, ymmword ptr [rsp + 32]
	vmovapd	ymmword ptr [rsp + 224], ymm1
	vmovapd	ymmword ptr [rsp + 192], ymm0
	vmovapd	ymm0, ymmword ptr [rsp + 224]
	vaddpd	ymm0, ymm0, ymmword ptr [rsp + 192]
	vmovapd	ymmword ptr [rsp + 96], ymm0
# %bb.3:                                #   in Loop: Header=BB0_1 Depth=1
	mov	eax, dword ptr [rsp + 92]
	add	eax, 4
	mov	dword ptr [rsp + 92], eax
	jmp	.LBB0_1
.LBB0_4:
	vmovapd	ymm0, ymmword ptr [rsp + 96]
	vmovapd	ymmword ptr [rsp + 288], ymm0
	vmovapd	ymmword ptr [rsp + 256], ymm0
	vmovapd	ymm0, ymmword ptr [rsp + 288]
	vmovapd	ymm1, ymmword ptr [rsp + 256]
	vhaddpd	ymm0, ymm0, ymm1
	vmovapd	ymmword ptr [rsp + 96], ymm0
	vmovapd	ymm0, ymmword ptr [rsp + 96]
	vmovapd	ymmword ptr [rsp + 320], ymm0
	vmovsd	xmm0, qword ptr [rsp + 320]     # xmm0 = mem[0],zero
	vmovapd	ymm1, ymmword ptr [rsp + 96]
	vextractf128	xmmword ptr [rsp + 368], ymm1, 1
	vmovsd	xmm1, qword ptr [rsp + 368]     # xmm1 = mem[0],zero
	vaddsd	xmm0, xmm0, xmm1
	mov	rax, qword ptr [rsp + 152]
	vmovsd	qword ptr [rax], xmm0
	mov	rax, qword ptr [rsp + 160]
	mov	eax, dword ptr [rax]
	sub	eax, dword ptr [rsp + 148]
	mov	dword ptr [rsp + 28], eax
.LBB0_5:                                # =>This Inner Loop Header: Depth=1
	mov	eax, dword ptr [rsp + 28]
	mov	rcx, qword ptr [rsp + 160]
	cmp	eax, dword ptr [rcx]
	jge	.LBB0_8
# %bb.6:                                #   in Loop: Header=BB0_5 Depth=1
	mov	rax, qword ptr [rsp + 168]
	movsxd	rcx, dword ptr [rsp + 28]
	vmovsd	xmm0, qword ptr [rax + 8*rcx]   # xmm0 = mem[0],zero
	mov	rax, qword ptr [rsp + 152]
	vaddsd	xmm0, xmm0, qword ptr [rax]
	vmovsd	qword ptr [rax], xmm0
# %bb.7:                                #   in Loop: Header=BB0_5 Depth=1
	mov	eax, dword ptr [rsp + 28]
	add	eax, 1
	mov	dword ptr [rsp + 28], eax
	jmp	.LBB0_5
.LBB0_8:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.Lfunc_end0:
	.size	sum_float64_avx_intrinsics, .Lfunc_end0-sum_float64_avx_intrinsics
                                        # -- End function
	.globl	sum_float64_c                   # -- Begin function sum_float64_c
	.p2align	4, 0x90
	.type	sum_float64_c,@function
sum_float64_c:                          # @sum_float64_c
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	sub	rsp, 32
	mov	qword ptr [rsp + 24], rdi
	mov	qword ptr [rsp + 16], rsi
	mov	qword ptr [rsp + 8], rdx
	mov	dword ptr [rsp + 4], 0
.LBB1_1:                                # =>This Inner Loop Header: Depth=1
	mov	eax, dword ptr [rsp + 4]
	mov	rcx, qword ptr [rsp + 16]
	cmp	eax, dword ptr [rcx]
	jge	.LBB1_4
# %bb.2:                                #   in Loop: Header=BB1_1 Depth=1
	mov	rax, qword ptr [rsp + 24]
	movsxd	rcx, dword ptr [rsp + 4]
	vmovsd	xmm0, qword ptr [rax + 8*rcx]   # xmm0 = mem[0],zero
	mov	rax, qword ptr [rsp + 8]
	vaddsd	xmm0, xmm0, qword ptr [rax]
	vmovsd	qword ptr [rax], xmm0
# %bb.3:                                #   in Loop: Header=BB1_1 Depth=1
	mov	eax, dword ptr [rsp + 4]
	add	eax, 1
	mov	dword ptr [rsp + 4], eax
	jmp	.LBB1_1
.LBB1_4:
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end1:
	.size	sum_float64_c, .Lfunc_end1-sum_float64_c
                                        # -- End function
	.globl	sum_grouped_float64_c           # -- Begin function sum_grouped_float64_c
	.p2align	4, 0x90
	.type	sum_grouped_float64_c,@function
sum_grouped_float64_c:                  # @sum_grouped_float64_c
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	sub	rsp, 48
	mov	qword ptr [rsp + 40], rdi
	mov	qword ptr [rsp + 32], rsi
	mov	qword ptr [rsp + 24], rdx
	mov	qword ptr [rsp + 16], rcx
	mov	qword ptr [rsp + 8], r8
	mov	dword ptr [rsp + 4], 0
.LBB2_1:                                # =>This Inner Loop Header: Depth=1
	mov	eax, dword ptr [rsp + 4]
	mov	rcx, qword ptr [rsp + 16]
	cmp	eax, dword ptr [rcx]
	jge	.LBB2_4
# %bb.2:                                #   in Loop: Header=BB2_1 Depth=1
	mov	rax, qword ptr [rsp + 24]
	movsxd	rcx, dword ptr [rsp + 4]
	mov	eax, dword ptr [rax + 4*rcx]
	mov	dword ptr [rsp], eax
	mov	rax, qword ptr [rsp + 40]
	movsxd	rcx, dword ptr [rsp]
	vmovsd	xmm0, qword ptr [rax + 8*rcx]   # xmm0 = mem[0],zero
	mov	rax, qword ptr [rsp + 8]
	vaddsd	xmm0, xmm0, qword ptr [rax]
	vmovsd	qword ptr [rax], xmm0
# %bb.3:                                #   in Loop: Header=BB2_1 Depth=1
	mov	eax, dword ptr [rsp + 4]
	add	eax, 1
	mov	dword ptr [rsp + 4], eax
	jmp	.LBB2_1
.LBB2_4:
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end2:
	.size	sum_grouped_float64_c, .Lfunc_end2-sum_grouped_float64_c
                                        # -- End function
	.ident	"Ubuntu clang version 14.0.0-1ubuntu1"
	.section	".note.GNU-stack","",@progbits
	.addrsig
