(Sys.init)
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
@THIS
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
@5000
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
@1
D=A
@THIS
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
@RETURN1
D=A
@THAT
D=M
@THIS
D=M
@ARG
D=M
@LCL
D=M
@5
D=A
@SP
D=M-D
@ARG
M=D
@SP
D=M
@LCL
M=D
Sys.main
0;JMP
(RETURN1)
@SP
M=M-1
A=M
D=M
@13
M=D
@1
D=A
@5
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
(LOOP)
@LOOP
0;JMP
(Sys.main)
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@4001
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
@THIS
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
@5001
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
@1
D=A
@THIS
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
@200
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
@1
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
@40
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
@2
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
@6
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
@3
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
@123
D=A
@SP
A=M
M=D
@SP
M=M+1
@RETURN2
D=A
@THAT
D=M
@THIS
D=M
@ARG
D=M
@LCL
D=M
@6
D=A
@SP
D=M-D
@ARG
M=D
@SP
D=M
@LCL
M=D
Sys.add12
0;JMP
(RETURN2)
@SP
M=M-1
A=M
D=M
@13
M=D
@0
D=A
@5
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
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
@1
D=A
@LOCAL
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@2
D=A
@LOCAL
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@3
D=A
@LOCAL
AD=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
@4
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
@SP
M=M-1
A=M
D=M+D
@SP
A=M
M=D
@SP
M=M+1
@5
D=A
@LCL
A=M-D
D=M
@15
M=D
@SP
M=M-1
A=M
@ARG
A=M
M=D
D=A+1
@SP
M=D
@LCL
AM=M-1
D=M
@THAT
M=D
@LCL
AM=M-1
D=M
@THIS
M=D
@LCL
AM=M-1
D=M
@ARG
M=D
@LCL
AM=M-1
D=M
@LCL
M=D
@15
A=M
0;JMP
(Sys.add12)
@4002
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
@THIS
AD=A+D
@14
M=D
@13
D=M
@14
A=M
M=D
@5002
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
@1
D=A
@THIS
AD=A+D
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
@12
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
D=M+D
@SP
A=M
M=D
@SP
M=M+1
@5
D=A
@LCL
A=M-D
D=M
@15
M=D
@SP
M=M-1
A=M
@ARG
A=M
M=D
D=A+1
@SP
M=D
@LCL
AM=M-1
D=M
@THAT
M=D
@LCL
AM=M-1
D=M
@THIS
M=D
@LCL
AM=M-1
D=M
@ARG
M=D
@LCL
AM=M-1
D=M
@LCL
M=D
@15
A=M
0;JMP
