/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated C++ implementation file in order to allow easy
development of Numbers library. It needs to be generated only once.

Interface version: 1.0.0

*/

#include "numbers_abi.hpp"
#include "numbers_interfaces.hpp"
#include "numbers_interfaceexception.hpp"

#include "numbers_variable.hpp"

using namespace Numbers;
using namespace Numbers::Impl;

IVariable * CWrapper::CreateVariable(const Numbers_double dInitialValue)
{
	return new CVariable(dInitialValue);
}

void CWrapper::GetVersion(Numbers_uint32 & nMajor, Numbers_uint32 & nMinor, Numbers_uint32 & nMicro)
{
	nMajor = NUMBERS_VERSION_MAJOR;
	nMinor = NUMBERS_VERSION_MINOR;
	nMicro = NUMBERS_VERSION_MICRO;
}

bool CWrapper::GetLastError(IBase* pInstance, std::string & sErrorMessage)
{
	if (pInstance) {
		return pInstance->GetLastErrorMessage(sErrorMessage);
	}
	else {
		throw ENumbersInterfaceException(NUMBERS_ERROR_INVALIDPARAM);
	}
}

void CWrapper::ReleaseInstance(IBase* pInstance)
{
	IBase::ReleaseBaseClassInterface(pInstance);
}

void CWrapper::AcquireInstance(IBase* pInstance)
{
	IBase::AcquireBaseClassInterface(pInstance);
}


