import os
import socket
try:
    from OpenSSL import crypto
except:
    pass

class essentials(object):
    def __init__(self):
        pass

    def enable_linux_iproute(self):
        file_path = "/proc/sys/net/ipv4/ip_forward"
        with open(file_path) as f:
            if f.read() == 1:
                return
        with open(file_path, "w") as f:
            print(1, file=f)

    def enable_windows_iproute(self):
        from services import WService
        service = WService("RemoteAccess")
        service.start()

    def enable_ip_route(self):
        runner.enable_windows_iproute() if "nt" in os.name else runner.enable_linux_iproute()

    def get_local_ip(self):
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        try:
            s.connect(('10.255.255.255', 1))
            IP = s.getsockname()[0]
        except Exception:
            IP = '127.0.0.1'
        finally:
            s.close()
        return IP

    def generate_ssl_cert(self, common_name, cert_file, key_file):
        k = crypto.PKey()
        k.generate_key(crypto.TYPE_RSA, 1024)

        cert = crypto.X509()
        cert.get_subject().C = "KE"
        cert.get_subject().ST = "Nairobi"
        cert.get_subject().L = "Langas"
        cert.get_subject().O = "Cyberafrics"
        cert.get_subject().OU = "Redteam"
        cert.get_subject().CN = common_name
        cert.set_serial_number(1000)
        cert.gmtime_adj_notBefore(0)
        cert.gmtime_adj_notAfter(10*365*24*60*60)
        cert.set_issuer(cert.get_subject())
        cert.set_pubkey(k)
        cert.sign(k, 'sha256')

        with open(cert_file, "wt") as f:
            f.write(crypto.dump_certificate(crypto.FILETYPE_PEM, cert).decode('utf-8'))
        with open(key_file, "wt") as f:
            f.write(crypto.dump_privatekey(crypto.FILETYPE_PEM, k).decode('utf-8'))

    def africana_gen_certs(self):
        if os.path.exists('/root/.africana/certs/'):
            pass
        elif not os.path.exists('/root/.africana/certs/'):
            os.system('mkdir -p /root/.africana/certs/')
            runner.generate_ssl_cert(common_name = 'africana', cert_file = '/root/.africana/certs/africana-cert.pem', key_file = '/root/.africana/certs/africana-key.pem')

    def kona_mbaya(self):
        runner.enable_ip_route()
        runner.africana_gen_certs()

runner = essentials()
