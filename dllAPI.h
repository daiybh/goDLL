#ifndef HELLO_H

#define HELLO_H

#ifdef __cplusplus
extern "C" {
#endif
	struct tagProviderHandle;
	//typedef void* ProviderHandle;
	typedef void* Frame;
	__declspec(dllexport) tagProviderHandle* openProvider(int providerID);
	__declspec(dllexport) Frame popFrame(tagProviderHandle* _ProviderHandle);
	__declspec(dllexport) bool pushVideoFrame(tagProviderHandle* _ProviderHandle, Frame frame, int nSize);
	__declspec(dllexport) bool pushAudioFrame(tagProviderHandle* _ProviderHandle, Frame frame, int nSize);

#ifdef __cplusplus
}
#endif

#endif