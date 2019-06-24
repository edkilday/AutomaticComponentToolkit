(*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated Pascal type definition file in order to allow easy
development of Numbers library. The functions in this file need to be implemented. It needs to be generated only once.

Interface version: 1.0.0

*)

{$MODE DELPHI}
unit numbers_types;

interface

uses
  Classes,
  sysutils;

(*************************************************************************************************************************
 Version definition for Numbers
**************************************************************************************************************************)

const
  NUMBERS_VERSION_MAJOR = 1;
  NUMBERS_VERSION_MINOR = 0;
  NUMBERS_VERSION_MICRO = 0;
  NUMBERS_VERSION_PRERELEASEINFO = '';
  NUMBERS_VERSION_BUILDINFO = '';


(*************************************************************************************************************************
 General type definitions
**************************************************************************************************************************)

type
  TNumbersResult = Cardinal;
  TNumbersHandle = Pointer;

  PNumbersResult = ^TNumbersResult;
  PNumbersHandle = ^TNumbersHandle;

(*************************************************************************************************************************
 Error Constants for Numbers
**************************************************************************************************************************)

const
  NUMBERS_SUCCESS = 0;
  NUMBERS_ERROR_NOTIMPLEMENTED = 1;
  NUMBERS_ERROR_INVALIDPARAM = 2;
  NUMBERS_ERROR_INVALIDCAST = 3;
  NUMBERS_ERROR_BUFFERTOOSMALL = 4;
  NUMBERS_ERROR_GENERICEXCEPTION = 5;
  NUMBERS_ERROR_COULDNOTLOADLIBRARY = 6;
  NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT = 7;
  NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION = 8;


implementation

end.
