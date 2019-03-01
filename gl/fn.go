// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import "unsafe"

type call struct {
	args     fnargs
	parg     unsafe.Pointer
	blocking bool
}

type fnargs struct {
	fn glfn

	a0 uintptr
	a1 uintptr
	a2 uintptr
	a3 uintptr
	a4 uintptr
	a5 uintptr
	a6 uintptr
	a7 uintptr
	a8 uintptr
	a9 uintptr
}

type glfn int

const (
	glfnUNDEFINED glfn = iota
	glfnActiveTexture
	glfnAttachShader
	glfnBindAttribLocation
	glfnBindBuffer
	glfnBindFramebuffer
	glfnBindRenderbuffer
	glfnBindTexture
	glfnBindVertexArray
	glfnBlendColor
	glfnBlendEquation
	glfnBlendEquationSeparate
	glfnBlendFunc
	glfnBlendFuncSeparate
	glfnBufferData
	glfnBufferSubData
	glfnCheckFramebufferStatus
	glfnClear
	glfnClearColor
	glfnClearDepthf
	glfnClearStencil
	glfnColorMask
	glfnCompileShader
	glfnCompressedTexImage2D
	glfnCompressedTexSubImage2D
	glfnCopyTexImage2D
	glfnCopyTexSubImage2D
	glfnCreateProgram
	glfnCreateShader
	glfnCullFace
	glfnDeleteBuffer
	glfnDeleteFramebuffer
	glfnDeleteProgram
	glfnDeleteRenderbuffer
	glfnDeleteShader
	glfnDeleteTexture
	glfnDeleteVertexArray
	glfnDepthFunc
	glfnDepthRangef
	glfnDepthMask
	glfnDetachShader
	glfnDisable
	glfnDisableVertexAttribArray
	glfnDrawArrays
	glfnDrawElements
	glfnEnable
	glfnEnableVertexAttribArray
	glfnFinish
	glfnFlush
	glfnFramebufferRenderbuffer
	glfnFramebufferTexture2D
	glfnFrontFace
	glfnGenBuffer
	glfnGenFramebuffer
	glfnGenRenderbuffer
	glfnGenTexture
	glfnGenVertexArray
	glfnGenerateMipmap
	glfnGetActiveAttrib
	glfnGetActiveUniform
	glfnGetAttachedShaders
	glfnGetAttribLocation
	glfnGetBooleanv
	glfnGetBufferParameteri
	glfnGetError
	glfnGetFloatv
	glfnGetFramebufferAttachmentParameteriv
	glfnGetIntegerv
	glfnGetProgramInfoLog
	glfnGetProgramiv
	glfnGetRenderbufferParameteriv
	glfnGetShaderInfoLog
	glfnGetShaderPrecisionFormat
	glfnGetShaderSource
	glfnGetShaderiv
	glfnGetString
	glfnGetTexParameterfv
	glfnGetTexParameteriv
	glfnGetUniformLocation
	glfnGetUniformfv
	glfnGetUniformiv
	glfnGetVertexAttribfv
	glfnGetVertexAttribiv
	glfnHint
	glfnIsBuffer
	glfnIsEnabled
	glfnIsFramebuffer
	glfnIsProgram
	glfnIsRenderbuffer
	glfnIsShader
	glfnIsTexture
	glfnLineWidth
	glfnLinkProgram
	glfnPixelStorei
	glfnPolygonOffset
	glfnReadPixels
	glfnReleaseShaderCompiler
	glfnRenderbufferStorage
	glfnSampleCoverage
	glfnScissor
	glfnShaderSource
	glfnStencilFunc
	glfnStencilFuncSeparate
	glfnStencilMask
	glfnStencilMaskSeparate
	glfnStencilOp
	glfnStencilOpSeparate
	glfnTexImage2D
	glfnTexParameterf
	glfnTexParameterfv
	glfnTexParameteri
	glfnTexParameteriv
	glfnTexSubImage2D
	glfnUniform1f
	glfnUniform1fv
	glfnUniform1fvP
	glfnUniform1fvUP
	glfnUniform1i
	glfnUniform1iv
	glfnUniform1ivP
	glfnUniform1ivUP
	glfnUniform2f
	glfnUniform2fv
	glfnUniform2fvP
	glfnUniform2fvUP
	glfnUniform2i
	glfnUniform2iv
	glfnUniform2ivP
	glfnUniform2ivUP
	glfnUniform3f
	glfnUniform3fv
	glfnUniform3fvP
	glfnUniform3fvUP
	glfnUniform3i
	glfnUniform3iv
	glfnUniform3ivP
	glfnUniform3ivUP
	glfnUniform4f
	glfnUniform4fv
	glfnUniform4fvP
	glfnUniform4fvUP
	glfnUniform4i
	glfnUniform4iv
	glfnUniform4ivP
	glfnUniform4ivUP
	glfnUniformMatrix2fv
	glfnUniformMatrix2fvP
	glfnUniformMatrix2fvUP
	glfnUniformMatrix3fv
	glfnUniformMatrix3fvP
	glfnUniformMatrix3fvUP
	glfnUniformMatrix4fv
	glfnUniformMatrix4fvP
	glfnUniformMatrix4fvUP
	glfnUseProgram
	glfnValidateProgram
	glfnVertexAttrib1f
	glfnVertexAttrib1fv
	glfnVertexAttrib2f
	glfnVertexAttrib2fv
	glfnVertexAttrib3f
	glfnVertexAttrib3fv
	glfnVertexAttrib4f
	glfnVertexAttrib4fv
	glfnVertexAttribPointer
	glfnViewport

	// ES 3.0 functions
	glfnUniformMatrix2x3fv
	glfnUniformMatrix3x2fv
	glfnUniformMatrix2x4fv
	glfnUniformMatrix4x2fv
	glfnUniformMatrix3x4fv
	glfnUniformMatrix4x3fv
	glfnBlitFramebuffer
	glfnUniform1ui
	glfnUniform2ui
	glfnUniform3ui
	glfnUniform4ui
	glfnUniform1uiv
	glfnUniform2uiv
	glfnUniform3uiv
	glfnUniform4uiv
)

func goString(buf []byte) string {
	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	panic("buf is not NUL-terminated")
}

func glBoolean(b bool) uintptr {
	if b {
		return TRUE
	}
	return FALSE
}
