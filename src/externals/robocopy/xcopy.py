#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import time
import shutil
import threading
from src.core.bcolors import *

FilesLeft = 0

def FullFolderSize(path):
    TotalSize = 0
    if os.path.exists(path):
        for root, dirs, files in os.walk(path):
            for file in files:
                TotalSize += os.path.getsize(os.path.join(root, file))
    return TotalSize

def getPERCECENTprogress(source_path, destination_path, bytes_to_copy):
    dstINIsize = FullFolderSize(destination_path)
    time.sleep(.25)
    print(" ")
    print(bcolors.BOLD + bcolors.GREEN + "FROM:" + bcolors.ENDC + "   ", source_path)
    print(bcolors.BOLD + bcolors.GREEN + "TO  :" + bcolors.ENDC + "   ", destination_path)
    print(" ")
    if os.path.exists(destination_path):
        while bytes_to_copy != (FullFolderSize(destination_path)-dstINIsize):
            sys.stdout.write('\r')
            percentagem = int((float((FullFolderSize(destination_path)-dstINIsize))/float(bytes_to_copy)) * 100)
            steps = int(percentagem/5)
            copiado = '{:,}'.format(int((FullFolderSize(destination_path)-dstINIsize)/1000000))
            sizzz = '{:,}'.format(int(bytes_to_copy/1000000))
            sys.stdout.write(("{:s} / {:s} Mb  ".format(copiado, sizzz)) +  (bcolors.BOLD + bcolors.GREEN + "{:29s}".format('|'*steps) + bcolors.ENDC) + ("  {:d}% ".format(percentagem)) + ("  {:d} ToGo \n".format(FilesLeft)))
            sys.stdout.flush()
            time.sleep(.01)
        sys.stdout.write('\r')
        time.sleep(.05)
        sys.stdout.write(("\n{:s} / {:s} Mb  ".format('{:,}'.format(int((FullFolderSize(destination_path)-dstINIsize)/1000000)), '{:,}'.format(int(bytes_to_copy/1000000)))) +  (bcolors.BOLD + bcolors.backGreen + "{:20s}".format(' '*29) + bcolors.ENDC) + ("   {:d}%   ".format( 100)) + ("{:s}   ".format('   ')) + "\n")
        sys.stdout.flush()
        print(" ")

def CopyProgress(SOURCE, DESTINATION):
    global FilesLeft
    DST = os.path.join(DESTINATION, os.path.basename(SOURCE))
    DST = DESTINATION
    if DST.startswith(SOURCE):
        print(" ")
        print(bcolors.BOLD + bcolors.UNDERL + 'Source folder can\'t be changed.' + bcolors.ENDC)
        print('Please check your target path...')
        print(" ")
        print(bcolors.BOLD + '        CANCELED' + bcolors.ENDC)
        print(" ")
        exit()

    Bytes2copy = 0
    for root, dirs, files in os.walk(SOURCE):
        dstDIR = root.replace(SOURCE, DST, 1)
        for filename in files:
            dstFILE = os.path.join(dstDIR, filename)
            if os.path.exists(dstFILE): continue 
            Bytes2copy += os.path.getsize(os.path.join(root, filename))
            FilesLeft += 1
    threading.Thread(name='progresso', target=getPERCECENTprogress, args=(SOURCE, DST, Bytes2copy)).start()
    for root, dirs, files in os.walk(SOURCE):
        dstDIR = root.replace(SOURCE, DST, 1)
        if not os.path.exists(dstDIR):
            os.makedirs(dstDIR)
        for filename in files:
            srcFILE = os.path.join(root, filename)
            dstFILE = os.path.join(dstDIR, filename)
            if os.path.exists(dstFILE): continue 
            head, tail = os.path.splitext(filename)
            count = -1
            year = int(time.strftime("%Y"))
            month = int(time.strftime("%m"))
            day = int(time.strftime("%d"))
            hour = int(time.strftime("%H"))
            minute = int(time.strftime("%M"))
            while os.path.exists(dstFILE):
                count += 1
                if count == 0:
                    dstFILE = os.path.join(dstDIR, '{:s}[{:d}.{:d}.{:d}]{:d}-{:d}{:s}'.format(head, year, month, day, hour, minute, tail))
                else:
                    dstFILE = os.path.join(dstDIR, '{:s}[{:d}.{:d}.{:d}]{:d}-{:d}[{:d}]{:s}'.format(head, year, month, day, hour, minute, count, tail))

            shutil.copy2(srcFILE, dstFILE)
            FilesLeft -= 1
'''
Ex.
CopyProgress('/path/to/SOURCE', '/path/to/DESTINATION')
'''
