/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated C++ Implementation file with the basic internal
 exception type in order to allow an easy use of Numbers library

Interface version: 1.0.0

*/


#include <string>

#include "numbers_interfaceexception.hpp"

/*************************************************************************************************************************
 Class ENumbersInterfaceException
**************************************************************************************************************************/
ENumbersInterfaceException::ENumbersInterfaceException(NumbersResult errorCode)
	: m_errorMessage("Numbers Error " + std::to_string (errorCode))
{
	m_errorCode = errorCode;
}

ENumbersInterfaceException::ENumbersInterfaceException(NumbersResult errorCode, std::string errorMessage)
	: m_errorMessage(errorMessage + " (" + std::to_string (errorCode) + ")")
{
	m_errorCode = errorCode;
}

NumbersResult ENumbersInterfaceException::getErrorCode ()
{
	return m_errorCode;
}

const char * ENumbersInterfaceException::what () const noexcept
{
	return m_errorMessage.c_str();
}

