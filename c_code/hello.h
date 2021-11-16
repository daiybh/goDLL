#ifndef HELLO_H

#define HELLO_H

#include <stdio.h>
#include <stdint.h>
void HelloFunc();

typedef struct buffer_data { //缓存数据的结构体
	uint8_t* ptr;
	size_t size;
} buffer_data;

extern "C" __declspec(dllexport) int add(int a, int b);
extern "C" __declspec(dllexport) int push(char*buf, int size);

#ifdef __cplusplus
extern "C" {
#endif
	typedef void* Foo;
	typedef void* Frame;
	__declspec(dllexport) Foo FooInit(void);
	__declspec(dllexport) void FooBar(Foo);
	__declspec(dllexport) Frame popFrame(Foo);
	__declspec(dllexport) void pushFrame(Foo , Frame , int );
	__declspec(dllexport) void FooFree(Foo);
#ifdef __cplusplus
}
#endif

#endif