import socket


def foo():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect(("127.0.0.1", 8080))
    # s.close()


if __name__ == "__main__":
    while True:
        foo()
