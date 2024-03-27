# Stacker
A tool for saving things to a stack to share between tmux sessions

## What is this and why

I have a need to "copy and paste" things between tmux panes all the time. A
standard clipboard could be used, but I don't always have a clipboard (old twm
machine for instance) and some times I want the value to persist, or even
stranger, I want one copy and multiple pastes. I tried fifo files, but the copy
app is held till the paste app runs and it is not one to many but one to one.
Also, fifo will not persist through a reboot.

This solution is to create a simple clipboard cache in a file.

## Usage

Create something:

	>echo thing | stacker

Peek at something:

	>stacker -peek
	thing

Read something, consuming it:

	>stacker
	thing

Show everything in a tab delimitated list:

	>stacker -ls
	0 thing 1
	1 thing 2

Update something, replacing:

	>echo wrong-thing | stacker ; echo Thing | stacker -update

Delete something

	>stacker -delete

## Web Usage

Use the `-http` flag to open on port 3333.

| Method | Path  | Description |
| ------ | ----- | ----------- |
|   GET  | /list | List stack items
|   GET  | /     | Return top item on stack
|  POST  | /rot  | Rotate Stack
|   PUT  | /     | Push value onto stack
| DELETE | /     | Pop the stack
|  HEAD  | /*    | Info on holdings


## Development

To test the application, use:

	>go test ./tests/...

To build this application, take care to name the output because "stacker" is a
directory:

	>go build -o bin/stacker

---
Copyright 2024 by thomas.cherry@gmail.com, all rights reserved.
