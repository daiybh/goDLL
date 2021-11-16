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

	uint8_t *m_pbuffer[10];
	int m_size = 0;
	int m_pos = 0;
	cxxFoo(int _a) :a(_a) {
		for (int i = 0; i < 10; i++)
			m_pbuffer[i] = new uint8_t[10];
	};
	~cxxFoo() {};
	void Bar();
	uint8_t* popFrame() {
		return m_pbuffer[m_pos++%10];
	}
	void pushFrame(uint8_t*frame,int size) {
		m_size = size;

		printf("\naaaaaaaa--%p-%d\n",frame,size);
		for (int i = 0; i < 10; i++)
		{
			printf("%x ", frame[i]);
		}
		printf("\n");
		
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

void pushFrame(Foo f, Frame frame,int size)
{
	cxxFoo* foo = (cxxFoo*)f;
	foo->pushFrame((uint8_t*)frame, size);
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
