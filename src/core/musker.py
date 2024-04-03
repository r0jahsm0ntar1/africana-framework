import os
import re
import subprocess
from src.core.bcolors import *

class ProxyConfig:
    def __init__(self):
        self.proxy = ''
        self.proxy_type = ''

    def prompt_for_proxy(self):
        self.proxy = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.BLUE + "Proxy format" + bcolors.ENDC + ":" + bcolors.YELLOW + "http://127.0.0.1:3129" + bcolors.GREEN + ")# " + bcolors.ENDC)
        self.determine_proxy_type()

    def determine_proxy_type(self):
        if re.match(r'^https?://', self.proxy):
            self.proxy_type = 'http'
        elif re.match(r'^socks4://', self.proxy):
            self.proxy_type = 'socks4'
        elif re.match(r'^socks5://', self.proxy):
            self.proxy_type = 'socks5'
        else:
            sys.exit(1)

    def set_proxies(self):
        if self.proxy_type == 'http':
            os.environ['http_proxy'] = self.proxy
            os.environ['https_proxy'] = self.proxy
        elif self.proxy_type in ['socks4', 'socks5']:
            os.environ['all_proxy'] = self.proxy
        else:
            sys.exit(1)

class CommandRunner:
    def __init__(self, proxy_config):
        self.proxy_config = proxy_config

    def run_command(self, command):
        self.proxy_config.set_proxies()

proxy_config = ProxyConfig()

if ' __name__' == '__main__':
        sys.exit(proxy_config())

