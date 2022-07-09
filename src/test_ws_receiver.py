import asyncio

import websockets


async def handler(websocket):
    while True:
        message = await websocket.recv()
        print(message)


async def main():
    async with websockets.serve(handler, "", 3333):
        await asyncio.Future()  # run forever


if __name__ == "__main__":
    print("starting websocket..")
    asyncio.run(main())


