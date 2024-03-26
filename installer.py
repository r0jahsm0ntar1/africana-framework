#! /usr/bin/python3
# coding=utf-8
import os
from src.core.system import *
while True:
    try:
        installer.update_system(); break
    except:
        os.system('clear')
        beauty.graphics(), scriptures.verses()
        break

