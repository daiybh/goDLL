#include <iostream>
#include "AES.h"
using namespace std;
#include "cpp-base64\base64.h"

#include "cryptopp/cryptlib.h"
#include "out\build\x64-Debug\vcpkg_installed\x86-windows\include\cryptopp\modes.h"
#include "out\build\x64-Debug\vcpkg_installed\x86-windows\include\cryptopp\aes.h"
#include "out\build\x64-Debug\vcpkg_installed\x86-windows\include\cryptopp\filters.h"
//#include "cryptopp/dll.h"
using CryptoPP::Exception;
using namespace CryptoPP;

class MyAES {
public:
	int min(int a, int b) {
		return (a > b) ? b : a;
	}
	MyAES(std::string skey = "")
	{
		if (!skey.empty())
		{
			for (int i = 0; i < min(skey.size(), AES::DEFAULT_KEYLENGTH); i += 2)
			{
				key[i] = short_to_hex(skey[i], skey[i + 1]);
			}
		}
		enc.SetKey(key, sizeof(key));

	}
	std::string encrypion(std::string plain)
	{
		StreamTransformationFilter encryptor(enc, nullptr);
		for (int i = 0; i < plain.size(); i++)
			encryptor.Put((unsigned char)plain[i]);
		encryptor.MessageEnd();

		size_t ready = encryptor.MaxRetrievable();

		string ci(ready, 0x00);
		encryptor.Get((unsigned char*)&ci[0], ci.size());

		string hexString;

		char b[10];
		for (int i = 0; i < ci.size(); i++)
		{
			sprintf_s(b, "%02X", ci[i] & 0xFF);
			hexString += b;
		}


		return hexString;
	}
private:

	ECB_Mode< AES >::Encryption enc;

	uint8_t key[AES::DEFAULT_KEYLENGTH] = { 0xAD,0xED,0x51,0x17,0x17,0x67,0x49,0x7B,0x94,0x36,0x8A,0xF5,0x10,0x50,0x75,0x4D };

	int char_to_hex(char a)
	{
		if ((a >= 'a' && a <= 'f'))
			return (a - 'a') + 10;
		else if (a >= 'A' && a <= 'F')
			return (a - 'A') + 10;
		else if (a >= '0' && a <= '9')
			return a - '0';
		return 0;
	};
	int short_to_hex(char a, char b)
	{
		return char_to_hex(a) * 16 + char_to_hex(b);
	}

};
int main()
{
  cout << "Aes dev" << endl;
  string jsonS = "{\"loginName\":\"xinyzhtc\",\"interfaceCode\": \"xyUploadData\"}";
  //¼ÓÃÜ...
  // Console.WriteLine(AesEncrypt("{\r\n" + 
  //"    \"loginName\":\"xinyzhtc\",\r\n" + 
  //"    \"interfaceCode\":\"xyUploadData\"\r\n" + 
  //"}","9v1lMn5H8LVUzdX6"));
  //½âÃÜ
  string aeskey = "9v1lMn5H8LVUzdX6";

  MyAES myAES(aeskey);

  std::string aaa = myAES.encrypion(jsonS);

  /*
  AES aes(128);
  unsigned int outlen;
  unsigned char*c = aes.EncryptECB((unsigned char*)jsonS.data(),jsonS.length(),(unsigned char*)aeskey.data(), outlen);

  std::string orig=std::string((char*)c);

  std::string encoded = base64_encode(reinterpret_cast<const unsigned char*>(orig.c_str()), orig.length());
  std::string decoded = base64_decode(encoded);
  */
  return 0;
}
