import subprocess
import sys
import os

def run_command(cmd):
    """Run a command and return output"""
    try:
        result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
        return result.stdout
    except Exception as e:
        return f"Error: {e}"

def create_assembly_dump(binary_file, output_file):
    """Create assembly dump using objdump"""
    print(f"Creating assembly dump from {binary_file}...")
    
    assembly = run_command(f"objdump -d {binary_file}")
    
    with open(output_file, 'w') as f:
        f.write(f"Assembly dump of: {binary_file}\n")
        f.write(f"Created: {os.popen('date').read().strip()}\n")
        f.write("=" * 80 + "\n\n")
        f.write(assembly)
    
    print(f"Assembly saved to: {output_file}")
    return len(assembly.split('\n'))

def create_binary_dump(binary_file, output_file):
    """Create binary dump with memory addresses"""
    print(f"Creating binary dump from {binary_file}...")
    
    with open(binary_file, 'rb') as f:
        f.seek(0, 2)
        file_size = f.tell()
        f.seek(0)
        
        print(f"Processing {file_size:,} bytes...")
        

        section_info = run_command(f"readelf -S {binary_file} | grep '\.text'")
        
        with open(output_file, 'w') as f_out:
            f_out.write(f"Binary machine code dump of: {binary_file}\n")
            f_out.write(f"Created: {os.popen('date').read().strip()}\n")
            f_out.write(f"File size: {file_size:,} bytes\n")
            if section_info:
                f_out.write(f".text section: {section_info}")
            f_out.write("=" * 80 + "\n\n")
            

            text_virt = 0x401000  
            text_file = 0x1000    
            
            bytes_written = 0
            chunk_size = 16  
            
            for i in range(0, file_size, chunk_size):
                f.seek(i)
                chunk = f.read(chunk_size)
                if not chunk:
                    break
                
                if i >= text_file:
                    virt_addr = (i - text_file) + text_virt
                else:
                    virt_addr = i  
                
                f_out.write(f"0x{virt_addr:08x}: ")
                
                for j, byte in enumerate(chunk):
                    f_out.write(f"{byte:08b} ")
                
                for j in range(chunk_size - len(chunk)):
                    f_out.write("         ")
                
                f_out.write("  ")
                for byte in chunk:
                    if 32 <= byte <= 126:
                        f_out.write(chr(byte))
                    else:
                        f_out.write(".")
                
                f_out.write("\n")
                bytes_written += len(chunk)
                
                if i % 65536 == 0 and i > 0:
                    print(f"  Processed {i:,} bytes...")
    
    print(f"Binary dump saved to: {output_file}")
    print(f"Total bytes written: {bytes_written:,}")
    return bytes_written

def main():
    if len(sys.argv) != 2:
        print("Usage: python3 extract_final.py <binary_file>")
        sys.exit(1)
    
    binary_file = sys.argv[1]
    
    if not os.path.exists(binary_file):
        print(f"Error: File '{binary_file}' not found!")
        sys.exit(1)
    
    assembly_file = "all_assembly.txt"
    binary_file_out = "all_machine_code_binary.txt"
    
    print("=" * 60)
    print(f"Processing: {binary_file}")
    print("=" * 60)
    
    asm_lines = create_assembly_dump(binary_file, assembly_file)
    
    print("-" * 60)
    
    bin_bytes = create_binary_dump(binary_file, binary_file_out)
    
    print("=" * 60)
    print("SUMMARY:")
    print(f"  Assembly file: {assembly_file} ({asm_lines} lines)")
    print(f"  Binary file: {binary_file_out} ({bin_bytes} bytes represented)")
    print("=" * 60)

if __name__ == "__main__":
    main()
