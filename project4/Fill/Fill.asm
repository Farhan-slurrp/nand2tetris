// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

(LOOP)
    // Put screen at RAM[0]
    @SCREEN
    D=A
    @R0
    M=D

    (KYBDCHECK)
        @KYBD
        D=M

        // Black if pressed
        @BLACK
        D;JGT

        // White if not pressed
        @WHITE
        D;JEQ

        @KYBDCHECK
        0;JMP

    (BLACK)
        @R1
        M=-1
        @CHANGE
        0;JMP

    (WHITE)
        @R1
        M=0
        @CHANGE
        0;JMP

    (CHANGE)
        // Load color to fill the screen with
        @R1
        D=M

        // Get pixel address and fill it with color from D
        @R0
        A=M
        M=D

        // Increment to the next pixel
        @R0
        D=M+1
        @KBD
        D=A-D	//KBD-SCREEN=A

        // Increment to the next pixel
        @R0
        M=M+1
        A=M

        // Restart if screen havent been filled completely
        @CHANGE
        D;JGT

    @LOOP
    0;JMP