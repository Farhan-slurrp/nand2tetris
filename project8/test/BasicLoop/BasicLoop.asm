@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@13
M=D
@0
D=A
@LOCAL
AD=M+D
@14
M=D
@13
D=M
@14
A=M
M=D
(LOOP)
@0
D=A
@ARG
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@LOCAL
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M+D
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@13
M=D
@0
D=A
@LOCAL
AD=M+D
@14
M=D
@13
D=M
@14
A=M
M=D
@0
D=A
@ARG
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@13
M=D
@0
D=A
@ARG
AD=M+D
@14
M=D
@13
D=M
@14
A=M
M=D
@0
D=A
@ARG
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
@LOOP
D;JNE
@0
D=A
@LOCAL
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
