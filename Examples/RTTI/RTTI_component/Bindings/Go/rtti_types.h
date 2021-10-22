/*++

Copyright (C) 2021 ADSK

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated plain C Header file with basic types in
order to allow an easy use of RTTI

Interface version: 1.0.0

*/

#ifndef __RTTI_TYPES_HEADER
#define __RTTI_TYPES_HEADER

#include <stdbool.h>

/*************************************************************************************************************************
 Scalar types definition
**************************************************************************************************************************/

#ifdef RTTI_USELEGACYINTEGERTYPES

typedef unsigned char RTTI_uint8;
typedef unsigned short RTTI_uint16 ;
typedef unsigned int RTTI_uint32;
typedef unsigned long long RTTI_uint64;
typedef char RTTI_int8;
typedef short RTTI_int16;
typedef int RTTI_int32;
typedef long long RTTI_int64;

#else // RTTI_USELEGACYINTEGERTYPES

#include <stdint.h>

typedef uint8_t RTTI_uint8;
typedef uint16_t RTTI_uint16;
typedef uint32_t RTTI_uint32;
typedef uint64_t RTTI_uint64;
typedef int8_t RTTI_int8;
typedef int16_t RTTI_int16;
typedef int32_t RTTI_int32;
typedef int64_t RTTI_int64 ;

#endif // RTTI_USELEGACYINTEGERTYPES

typedef float RTTI_single;
typedef double RTTI_double;

/*************************************************************************************************************************
 General type definitions
**************************************************************************************************************************/

typedef RTTI_int32 RTTIResult;
#pragma pack (1)
typedef struct {
  void * Handle;
  RTTI_uint64 ClassTypeId;
} RTTIHandle;
#pragma pack ()
#define RTTIHandleNull { nullptr, 0 }
typedef void * RTTI_pvoid;

/*************************************************************************************************************************
 Version for RTTI
**************************************************************************************************************************/

#define RTTI_VERSION_MAJOR 1
#define RTTI_VERSION_MINOR 0
#define RTTI_VERSION_MICRO 0
#define RTTI_VERSION_PRERELEASEINFO ""
#define RTTI_VERSION_BUILDINFO ""

/*************************************************************************************************************************
 Error constants for RTTI
**************************************************************************************************************************/

#define RTTI_SUCCESS 0
#define RTTI_ERROR_NOTIMPLEMENTED 1 /** functionality not implemented */
#define RTTI_ERROR_INVALIDPARAM 2 /** an invalid parameter was passed */
#define RTTI_ERROR_INVALIDCAST 3 /** a type cast failed */
#define RTTI_ERROR_BUFFERTOOSMALL 4 /** a provided buffer is too small */
#define RTTI_ERROR_GENERICEXCEPTION 5 /** a generic exception occurred */
#define RTTI_ERROR_COULDNOTLOADLIBRARY 6 /** the library could not be loaded */
#define RTTI_ERROR_COULDNOTFINDLIBRARYEXPORT 7 /** a required exported symbol could not be found in the library */
#define RTTI_ERROR_INCOMPATIBLEBINARYVERSION 8 /** the version of the binary interface does not match the bindings interface */

/*************************************************************************************************************************
 Error strings for RTTI
**************************************************************************************************************************/

inline const char * RTTI_GETERRORSTRING (RTTIResult nErrorCode) {
  switch (nErrorCode) {
    case RTTI_SUCCESS: return "no error";
    case RTTI_ERROR_NOTIMPLEMENTED: return "functionality not implemented";
    case RTTI_ERROR_INVALIDPARAM: return "an invalid parameter was passed";
    case RTTI_ERROR_INVALIDCAST: return "a type cast failed";
    case RTTI_ERROR_BUFFERTOOSMALL: return "a provided buffer is too small";
    case RTTI_ERROR_GENERICEXCEPTION: return "a generic exception occurred";
    case RTTI_ERROR_COULDNOTLOADLIBRARY: return "the library could not be loaded";
    case RTTI_ERROR_COULDNOTFINDLIBRARYEXPORT: return "a required exported symbol could not be found in the library";
    case RTTI_ERROR_INCOMPATIBLEBINARYVERSION: return "the version of the binary interface does not match the bindings interface";
    default: return "unknown error";
  }
}

/*************************************************************************************************************************
 Declaration of handle classes 
**************************************************************************************************************************/

typedef RTTIHandle RTTI_Base;
typedef RTTIHandle RTTI_Animal;
typedef RTTIHandle RTTI_Mammal;
typedef RTTIHandle RTTI_Reptile;
typedef RTTIHandle RTTI_Giraffe;
typedef RTTIHandle RTTI_Tiger;
typedef RTTIHandle RTTI_Snake;
typedef RTTIHandle RTTI_Turtle;
typedef RTTIHandle RTTI_AnimalIterator;
typedef RTTIHandle RTTI_Zoo;

/*************************************************************************************************************************
 Declaration of structs
**************************************************************************************************************************/

#pragma pack (1)

typedef struct {
    RTTI_int32 m_X;
    RTTI_int32 m_Y;
} sRTTITestStruct;

#pragma pack ()


#endif // __RTTI_TYPES_HEADER
