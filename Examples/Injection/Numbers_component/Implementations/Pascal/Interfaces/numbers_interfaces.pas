(*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated Pascal interface definition file in order to allow easy
development of Numbers library. The functions in this file need to be implemented. It needs to be generated only once.

Interface version: 1.0.0

*)

{$MODE DELPHI}
{$INTERFACES CORBA}
unit numbers_interfaces;

interface

uses
  numbers_types,
  Classes,
  sysutils;


type

(*************************************************************************************************************************
 Interface definition for Base
**************************************************************************************************************************)

INumbersBase = interface
  ['{52FDFC07-2182-454F-963F-5F0F9A621D72}']

  function GetLastErrorMessage(out AErrorMessage: String): Boolean;
  procedure ClearErrorMessages();
  procedure RegisterErrorMessage(const AErrorMessage: String);
  procedure IncRefCount();
  function DecRefCount(): Boolean;
end;


(*************************************************************************************************************************
 Interface definition for Variable
**************************************************************************************************************************)

INumbersVariable = interface(INumbersBase)
  ['{9566C74D-1003-4C4D-BBBB-0407D1E2C649}']

  function GetValue(): Double;
  procedure SetValue(const AValue: Double);
end;

implementation

end.
