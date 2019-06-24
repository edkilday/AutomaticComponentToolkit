/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated C++ implementation file in order to allow easy
development of Prime Numbers Library. It needs to be generated only once.

Interface version: 1.2.0

*/

#include "libprimes_abi.hpp"
#include "libprimes_interfaces.hpp"
#include "libprimes_interfaceexception.hpp"

using namespace LibPrimes;
using namespace LibPrimes::Impl;

void CWrapper::GetVersion(LibPrimes_uint32 & nMajor, LibPrimes_uint32 & nMinor, LibPrimes_uint32 & nMicro)
{
	nMajor = LIBPRIMES_VERSION_MAJOR;
	nMinor = LIBPRIMES_VERSION_MINOR;
	nMicro = LIBPRIMES_VERSION_MICRO;
}

bool CWrapper::GetLastError(IBase* pInstance, std::string & sErrorMessage)
{
	throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

void CWrapper::ReleaseInstance(IBase* pInstance)
{
	pInstance->DecRefCount();
}

IFactorizationCalculator * CWrapper::CreateFactorizationCalculator()
{
	throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

ISieveCalculator * CWrapper::CreateSieveCalculator()
{
	throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NOTIMPLEMENTED);
}


