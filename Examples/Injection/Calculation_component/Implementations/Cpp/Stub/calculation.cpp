/*++

Copyright (C) 2019 Calculation developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated C++ implementation file in order to allow easy
development of Calculation library. It needs to be generated only once.

Interface version: 1.0.0

*/

#include "calculation_abi.hpp"
#include "calculation_interfaces.hpp"
#include "calculation_interfaceexception.hpp"

#include "calculation_calculator.hpp"

using namespace Calculation;
using namespace Calculation::Impl;

// Injected Components
Numbers::PWrapper CWrapper::sPNumbersWrapper;

ICalculator * CWrapper::CreateCalculator()
{
	return new CCalculator();
}

void CWrapper::GetVersion(Calculation_uint32 & nMajor, Calculation_uint32 & nMinor, Calculation_uint32 & nMicro)
{
	nMajor = CALCULATION_VERSION_MAJOR;
	nMinor = CALCULATION_VERSION_MINOR;
	nMicro = CALCULATION_VERSION_MICRO;
}

bool CWrapper::GetLastError(IBase* pInstance, std::string & sErrorMessage)
{
	if (pInstance) {
		return pInstance->GetLastErrorMessage(sErrorMessage);
	}
	else {
		throw ECalculationInterfaceException(CALCULATION_ERROR_INVALIDPARAM);
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


