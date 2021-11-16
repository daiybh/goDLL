#include "hello.h"
#include <iostream>
using namespace std;

void HelloFunc()
{

    cout << "Hello World/n";
}
int add(int a, int b)
{
    return a + b;
}
int push(char* buf, int size)
{
    printf("c size=%d", size);
    return size;
}
// foo.hpp
class cxxFoo {
public:
	int a;
	uint8_t m_buffer[1024];
	cxxFoo(int _a) :a(_a) {};
	~cxxFoo() {};
	void Bar();
	uint8_t* popFrame() {
		return m_buffer;
	}
	void pushFrame(uint8_t*frame,int size) {

	}
};

// foo.cpp
#include <iostream>
void
cxxFoo::Bar(void) {
	std::cout << "c bar  "<<this->a << std::endl;
}

Foo FooInit()
{
	cxxFoo* ret = new cxxFoo(1);
	return (void*)ret;
}

void pushFrame(Foo, Frame)
{

}

void FooFree(Foo f)
{
	cxxFoo* foo = (cxxFoo*)f;
	delete foo;
}
void FooBar(Foo f)
{
	cxxFoo* foo = (cxxFoo*)f;
	foo->Bar();
}

Frame popFrame(Foo f)
{
	cxxFoo* foo = (cxxFoo*)f;
	return foo->popFrame();
}
