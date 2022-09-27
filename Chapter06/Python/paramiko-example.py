import paramiko 

TARGET = {
        "hostname": "10.0.4.1",
        "username": "netlab",
        "password": "netlab",
}

ssh = paramiko.SSHClient() 
ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy()) 
ssh.connect(**TARGET) 
stdin, stdout, stderr = ssh.exec_command("uptime") 
stdin.close()
 
print(stdout.read().decode("ascii")) 
