.text
.globl _start
_start:
    movq $input, %r8  # Current string address
    movq $0, %r15     # Result value
    jmp reset_numbers

loop:
    # Skip character if less than ascii '0'
    cmpb $0x30, (%r8)
    jl inc

    # Skip character if greater than ascii '9'
    cmpb $0x39, (%r8)
    jg inc

    movb (%r8), %r10b
    subb $0x30, %r10b

    # If first digit hasn't been found, update both
    cmpb $0x0, %r13b
    jne last_number

    movb %r10b, %r11b
    incb %r13b

last_number:
    movb %r10b, %r12b

inc:
    # TODO: next step is to check for a newline when incrementing,
    # if the value is a newline, then multiply first digit by 10,
    # add to result, then add second digit to result. Then reset numbers.
    incq %r8
    cmpq $input_end, %r8
    jge print_result

    cmpb $0x0a, (%r8)
    je add_to_result

    jmp loop

reset_numbers:
    movb $0, %r11b # first digit
    movb $0, %r12b # last digit
    movb $0, %r13b # first digit found flag
    jmp loop

add_to_result:
    imulq $10, %r11
    addq %r11, %r15
    addq %r12, %r15
    incq %r8
    cmpq $input_end, %r8
    jl reset_numbers

print_result:
    movq $result_end, %r14 # end of output area

    movq %r15, %rax
    movq $10, %rcx

print_loop:
    divq %rcx

    addq $0x30, %rdx
    movb %dl, (%r14)

    cmpq $0, %rax
    je print

    decq %r14
    movq $0, %rdx
    jmp print_loop

print:
    movq $0x04, %rax
    movq $0x1, %rbx
    movq %r14, %rcx
    movq $result_end, %rdx
    sub %r14, %rdx
    incq %rdx
    int $0x80

end:
    movl $0x01, %eax
    movl %r15d, %ebx
    int $0x80

.data
result: .ascii "         "
result_end: result_len = . - result
input: .ascii "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
"
input_end: input_len = . - input
