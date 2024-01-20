.text
.globl _start
_start:
    movq $input, %r8  # Address of current character of input string is stored in %r8
    movq $0, %r15     # Result value is stored in %r15
    jmp reset_numbers # Jump to pre-loop register reset, to initialize first & last digits and the first-digit-found flag

reset_numbers:
    movb $0, %r11b # first digit
    movb $0, %r12b # last digit
    movb $0, %r13b # first digit found flag

loop:
    cmpb $0x30, (%r8) # Skip character if less than ascii '0'
    jl inc

    cmpb $0x39, (%r8) # Skip character if greater than ascii '9'
    jg inc

    movb (%r8), %r10b # Copy current character into r10
    subb $0x30, %r10b # convert from ascii to plain integer

    cmpb $0x0, %r13b  # If the first digit has been found
    jne last_digit    # jump directly to updating the last digit

    movb %r10b, %r11b # Update the first number for the line (%r11b)
    incb %r13b        # and mark the first number as found (%r13b)

last_digit:
    movb %r10b, %r12b    # Update the last digit found

inc:
    incq %r8             # Move character pointer
    cmpq $input_end, %r8 # Check if we're at the end of the input
    jge print_result     # and jump to printing the result if so

    cmpb $0x0a, (%r8)    # Check if the current character is a newline,
    je add_to_result     # and add to final result if so

    jmp loop             # Jump to beginning of the main loop

add_to_result:
    imulq $10, %r11      # Multiply first digit by 10
    addq %r11, %r15      # Add first digit to result
    addq %r12, %r15      # Add last digit to result
    incq %r8             # Move character pointer forward
    cmpq $input_end, %r8 # Check if we're at the end of the input
    jl reset_numbers     # and jump to printing the result if so

print_result:
    movq $result_end, %r14 # Keep output pointer (end of output location) in %r14

    movq %r15, %rax  # Store result (dividend) in %rax
    movq $10, %rcx   # Store divisor in %rcx

print_loop:
    movq $0, %rdx    # Reset first 64 bits of dividend to 0
    divq %rcx        # Divide!

    addq $0x30, %rdx # Remainder is stored in %rdx, so add 0x30 to make it ascii
    movb %dl, (%r14) # Put character at output pointer location

    cmpq $0, %rax    # If the result of division was 0
    je print         # jump to print the result

    decq %r14        # Move character pointer left
    jmp print_loop   # Go back to the start of the loop

print:
    movq $0x04, %rax       # 4 is write()
    movq $0x1, %rbx        # File descriptor 1 (stdout)
    movq %r14, %rcx        # Pointer to result string
    movq $result_end, %rdx # Store address of result string as output length
    sub %r14, %rdx         # and subtract the beginning of the string
    incq %rdx              # and add 1 to get the actual length of the string
    int $0x80              # Send the syscall interrupt

end:
    movl $1, %eax # 1 is exit()
    movl $0, %ebx # Exit status 0
    int $0x80     # Send the syscall interrupt

.data
result: .ascii "         "
result_end:
input: .ascii "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
"
input_end:
