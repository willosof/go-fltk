//go:build darwin && arm64

package fltk

// #cgo darwin,arm64 CXXFLAGS: -std=c++11
// #cgo darwin,arm64 CPPFLAGS: -I${SRCDIR}/lib/darwin/arm64 -I${SRCDIR}/include -I${SRCDIR}/include/FL/images -I${SRCDIR}/include/png -I${SRCDIR}/include/zlib -I${SRCDIR}/include/jpeg -isysroot /Library/Developer/CommandLineTools/SDKs/MacOSX15.0.sdk -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo darwin,arm64 LDFLAGS: ${SRCDIR}/lib/darwin/arm64/libfltk_images.a ${SRCDIR}/lib/darwin/arm64/libfltk_jpeg.a ${SRCDIR}/lib/darwin/arm64/libfltk_png.a ${SRCDIR}/lib/darwin/arm64/libfltk_z.a ${SRCDIR}/lib/darwin/arm64/libfltk_gl.a -framework OpenGL ${SRCDIR}/lib/darwin/arm64/libfltk_forms.a ${SRCDIR}/lib/darwin/arm64/libfltk.a -lm -lpthread -framework Cocoa
import "C"
