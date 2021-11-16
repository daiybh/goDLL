cmake_minimum_required(VERSION 3.15)
	

set(SDKTOOLS_VERSION 1.0.0)
project(WU_Charging20211022
	VERSION ${SDKTOOLS_VERSION}
	DESCRIPTION "A simple utf-8 based unicode decoding and encoding library"
	HOMEPAGE_URL "https://github.com/northwindtraders/beicode"
	LANGUAGES CXX)

set(CMAKE_CXX20_STANDARD_COMPILE_OPTION "-std:c++latest")
set(CMAKE_CXX20_EXTENSION_COMPILE_OPTION "-std:c++latest")

if(CMAKE_BUILD_TYPE MATCHES Debug)
set(installpath Debug)
else()
set(installpath Release)
endif()

set(CMAKE_INSTALL_PREFIX ${CMAKE_SOURCE_DIR}/../  CACHE PATH "Installation Directory" FORCE)    

#add_definitions(-DUNICODE -D_UNICODE)
if(DEFINED ENV{VCPKG_ROOT} )
	if(NOT DEFINED CMAKE_TOOLCHAIN_FILE)
		set(CMAKE_TOOLCHAIN_FILE "$ENV{VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake" CACHE STRING "")
		else()
		
	Message("NOT DEFINED CMAKE_TOOLCHAIN_FILE")
	endif()
else()
	Message("DEFINED ENV{VCPKG_ROOT}")
endif()

SET (LIBHELLO_SRC hello.cpp)
ADD_LIBRARY (helloDLL SHARED ${LIBHELLO_SRC})
#ADD_LIBRARY (hello_static STATIC ${LIBHELLO_SRC})

add_subdirectory(test)
#add_executable(AAA 	"test.c" "pkcs7_padding.c" "AES.c")
INSTALL (TARGETS helloDLL  LIBRARY DESTINATION lib
RUNTIME DESTINATION bin
ARCHIVE DESTINATION lib)

INSTALL (FILES hello.h DESTINATION include/hello)

 