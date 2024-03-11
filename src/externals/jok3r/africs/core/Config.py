#!/usr/bin/env python3
# -*- coding: utf-8 -*-
###
### Core > Config
###
import colored
import os

from lib.core.Constants import *
from lib._version import __version__

#----------------------------------------------------------------------------------------
# Banner/Help

BANNER = colored.stylize("""
                     .` 
                   /dMMs`
                `oNMMMMMN+
              -yMMMMMMMMMMm/""", colored.attr('bold')) + colored.stylize("""
            :hMMMMMMMMMMMMMMd:""", colored.attr('bold')) + colored.stylize("""
         `+mMMMMho:sNMMMMMMMMMh.""", colored.attr('bold')) + colored.stylize("""
       .sNMMdo:     `/dMMMMMMMMMy`""", colored.attr('bold')) + colored.stylize("""
     -yMms/`           -sNMMMMMMMNo`""", colored.attr('bold')) + colored.stylize("""
    /y/.                 `+mMMMMMMMN+""", colored.attr('bold')) + colored.stylize("""
                            -yMMMMMMMm/
             """, colored.attr('bold')) + \
colored.stylize("..", colored.fg('light_green')+colored.attr('bold')) + colored.stylize(
  """               `+mMMMMMMd-
             """, colored.attr('bold')) + \
colored.stylize("-s+-`", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""               :yMMMMMMh.
             """, colored.attr('bold')) + \
colored.stylize(".ssso/-`", colored.fg('light_green')+colored.attr('bold')) + colored.stylize(
  """              `omMMMMMs`             """, colored.attr('bold')) + \
colored.stylize(  """                 
              """, colored.attr('bold')) + \
colored.stylize("ossssso/-`", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""              :hMMMMNo`
              """, colored.attr('bold')) + \
colored.stylize("/ssssss/.::-.", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""             .oNMMMN/
              """, colored.attr('bold')) + \
colored.stylize("`sssssso-   -:/:.", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""            /hMMMm:
               """, colored.attr('bold')) + \
colored.stylize("/ssssss+    :/sso+/-`", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""         .sNMMdsss:
               """, colored.attr('bold')) + \
colored.stylize("`ossssss-  : +ssssssso+/:.`", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""      /dMNs.
                 """, colored.attr('bold')) + \
colored.stylize("`.-:/++:/ .ossssssssssssso+/:-.`", colored.fg('light_green')+colored.attr('bold')) + colored.stylize(""" .`
                       """, colored.attr('bold')) + \
colored.stylize("oyso+-::://////:::--..``", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""
                       -MMMd
                        sMMN
                         dMM.
                         `mMo
                          `dN
                           `ho
                             o-

                 Rojahs Montari (c) 2024
             https://www.jok3r-framework.com
    https://github.com/r0jahsm0ntar1/africana-framework

[ """, colored.attr('bold')) + colored.stylize("      Network & Web Pentest Automation Framework", colored.fg('light_green')+colored.attr('bold')) + colored.stylize("""        ]""".format(version=__version__), colored.attr('bold'))

USAGE = """
python3 jok3r.py <command> [<args>]

Supported commands:
   toolbox    Manage the toolbox
   info       View supported services/options/checks
   db         Define missions scopes, keep tracks of targets & view attacks results
   attack     Run security checks against targets
"""

ATTACK_EXAMPLES = colored.stylize('Examples:', colored.attr('bold')) + """
  - Run all security checks against an URL in interactive mode (break before each check):
  python3 jok3r.py attack -t http://www.example.com/ 

  - Run all security checks against a MS-SQL service (without user interaction) and add results to the mission "mayhem" in db:
  python3 jok3r.py attack -t 192.168.1.42:1433 -s mssql --add2db mayhem --fast

  - Run only "recon" and "vulnscan" security checks against an FTP service and add results to the mission "mayhem" in db:
  python3 jok3r.py attack -t 192.168.1.142:21 -s ftp --cat-only recon,vulnscan --add2db mayhem

  - Run security checks against all FTP services running on 2121/tcp and all HTTP services from the mission "mayhem" in db:
  python3 jok3r.py attack -m mayhem -f "port=2121;service=ftp" -f "service=http"

  - Search for "easy wins" (critical vulns / easy to exploit) on all services registered in mission "mayhem" in db:
  python3 jok3r.py attack -m mayhem --profile red-team --fast
"""

DB_INTRO = """[                                                         ]
[ The local database stores the missions, targets info &  ]
[ attacks results.This shell allows for easy access to    ]
[ this database. New missions can be added and scopes can ]
[ be defined by importing new targets.                    ]
"""

#----------------------------------------------------------------------------------------
# Arguments Parsing Settings

ARGPARSE_MAX_HELP_POS    = 45
TARGET_FILTERS           = {
    'ip'       : FilterData.IP, 
    'host'     : FilterData.HOST,
    'port'     : FilterData.PORT, 
    'service'  : FilterData.SERVICE, 
    'url'      : FilterData.URL,
    'osfamily' : FilterData.OS_FAMILY,
    'banner'   : FilterData.BANNER,
}


#----------------------------------------------------------------------------------------
# Basic Settings

TOOL_BASEPATH      = os.path.dirname(os.path.realpath(__file__+'/../..'))
TOOLBOX_DIR        = TOOL_BASEPATH + '/toolbox'
DEFAULT_OUTPUT_DIR = 'output'
WEBSHELLS_DIR      = TOOL_BASEPATH + '/webshells'
WORDLISTS_DIR      = TOOL_BASEPATH + '/wordlists'
DB_FILE            = TOOL_BASEPATH + '/local.db'
DB_HIST_FILE       = TOOL_BASEPATH + '/.dbhistory'
REPORT_TPL_DIR     = TOOL_BASEPATH + '/lib/reporter/templates'
REPORT_PATH        = TOOL_BASEPATH + '/reports'
VIRTUALENVS_DIR    = TOOL_BASEPATH + '/toolbox/virtualenvs'

#----------------------------------------------------------------------------------------
# Display Settings

ATTACK_SUMMARY_TABLE_MAX_SIZE = 18

#----------------------------------------------------------------------------------------
# Settings Files

SETTINGS_DIR              = TOOL_BASEPATH + '/settings'
CONF_EXT                  = '.conf'
TOOLBOX_CONF_FILE         = 'toolbox'
INSTALL_STATUS_CONF_FILE  = '_install_status'
ATTACK_PROFILES_CONF_FILE = 'attack_profiles'
PREFIX_SECTION_CHECK      = 'check_'
MULTI_CONF                = 'multi'
MULTI_TOOLBOX_SUBDIR      = 'multi'

TOOL_OPTIONS = {
    MANDATORY: [
        'name',
        'description',
        'target_service',
    ],
    OPTIONAL: [
        'virtualenv',
        'install',
        'update',
        'check_command',
    ]
}

SERVICE_CHECKS_CONFIG_OPTIONS = {
    MANDATORY: [
        'default_port',
        'protocol',
        'categories',
    ],
    OPTIONAL: [
        'auth_types'
    ]
}

CHECK_OPTIONS = {
    MANDATORY: [
        'name',
        'category',
        'description',
        'tool',
        # command
    ],
    OPTIONAL: [
        'apikey',
    ]
}

OPTIONS_ENCRYTPED_PROTO = (
    'ftps',
    'https',
    'rmissl',
    'smtps',
    'telnets',
)

