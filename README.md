# Go Alone

*Go Alone* is a brief experiment into running Go as an
appliance-oriented operating system.

We do **not** run Go on bare metal, nor did we write a kernel in Go --
we use Linux, and just replace userspace with a Go app. *All of the
userspace.*

The single Go application just becames the `init`, PID 1, of the
machine. There is nothing really difficult here, it just works.

The scripts included build a minimal kernel, intended to run inside a
KVM. They don't even include support for block devices. Take that,
12-factor apps!

Here's how to run them:

    # if you have a local clone of linux.git, put it in ~/src/linux/
	# to speed up first run
	./do-build-kernel

	# build and run the hello world example app
	./do-run-hello-word

	# build and run the network app and forward port 8000 to it
	./do-run-network
	# in a browser, open http://localhost:8000/
	# when done, type control-A x

This assumes `x86_64` platform and a working install of
`qemu-system-x86_64`.

The output should look like this:

	...
    Unpacking initramfs...
    Freeing initrd memory: 580K (ffff880007f5f000 - ffff880007ff0000)
    io scheduler noop registered (default)
    ACPI: PCI Interrupt Link [LNKC] enabled at IRQ 11
    Serial: 8250/16550 driver, 4 ports, IRQ sharing disabled
    00:05: ttyS0 at I/O 0x3f8 (irq = 4, base_baud = 115200) is a 16550A
    Freeing unused kernel memory: 648K (ffffffff81279000 - ffffffff8131b000)
    Hello, world! I have 4 CPUs
    ACPI: Preparing to enter system sleep state S5
    reboot: Power down

## Roadmap

- play more with network configuration
- maybe set up virtfs
- maybe bundle asset files and/or config in the initramfs
- migrate from ad hoc scripts to utility commands, with
  configurability such as RAM size.. open questions:
  - what to do with generated files that are needed?
  - where to store them?
  - cache invisibly or make that part of the UI?
  - kernel can be shared between many apps, initramfs is per app (version)
- quick cli ideas

  ``` sh
  # build kernel
  alone build-kernel [--version=VER] [--mirror=PATH] OUT_PATH

  # create initramfs
  alone take INITRAMFS_PATH_TO_CREATE GO_PACKAGE

  # run the previously-created kernel and initramfs
  alone exec [--ram=512M] KERNEL INITRAMFS

  # sort of like `go run`, but takes package. builds kernel if needed
  alone run [--kernel-version=VER] [--ram=512M] GO_PACKAGE
  ```
