import asyncio, asyncssh, sys 
 
TARGET = { 
        "host": "10.0.4.1", 
        "username": "user", 
        "password": "pw", 
        "known_hosts": None,
} 
 
async def run_client() -> None: 
    async with asyncssh.connect(**TARGET) as conn: 
       result = await conn.run("uptime", check=True) 
       print(result.stdout, end="") 
 
try: 
    asyncio.get_event_loop().run_until_complete( run_client() ) 
except (OSError, asyncssh.Error) as execrr: 
        sys.exit("Connection failed:" + str(execrr)) 
