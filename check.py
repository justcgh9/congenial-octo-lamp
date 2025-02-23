import os
import subprocess
import time

def process_files(folder_path, binary_path, binary_args=[], go_binary_path=""):
    # print(os.listdir(folder_path))
    # while True:
    for filename in os.listdir(folder_path):
        file_path = os.path.join(folder_path, filename)
        # print(f"got to {file_path}")
        if os.path.isfile(file_path):
            with open(file_path, 'r') as file:
                # print(f"opened {file_path}")
                process = subprocess.Popen([binary_path] + binary_args, stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
                
                for line in file:
                    process.stdin.write(line)
                
                process.stdin.close()  # Send EOF
                errors = []
                for output_line in process.stdout:
                    if output_line.startswith("Type Error Tag: ["):
                        # print(f"{filename}: {output_line.strip()[17:-1]}")
                        errors.append(output_line.strip()[17:-1])
                
                go_process = subprocess.Popen([go_binary_path] + [file_path], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
                # print(errors)
                with open(file_path, 'r') as newfile:
                    for line in newfile:
                        go_process.stdin.write(line)
                
                go_process.stdin.close()
                output = []
                
                flag = False
                
                for output_line in go_process.stdout:
                    if output_line.startswith("ERROR"):
                        flag = True
                        # print(output_line.split(".")[0])
                        if output_line.split(".")[0].strip() not in errors:
                            print(f"{filename}: Expected one of {errors}, got {output_line.split(".")[0]}")
                
                if not flag and errors:
                    print(f"{filename}: Expected one of {errors}, got []")
                
                for line in go_process.stderr:
                    output.append(line)
                
                go_process.wait()
                status = go_process.returncode
                if status > 1:
                    print(f"Test: {filename}")
                    # for line in output:
                    #     print(line)
                # # elif errors:
                #     print(f"Test: {filename}")
                #     for line in errors:
                #         print(f'this error was not found {line}')
                
                # os.remove(file_path) 
        
        # time.sleep(1)  # Avoid tight loops, adjust as needed

if __name__ == "__main__":
    # folder_path = "./examples/public-tests/week-1/main/public/"
    # folder_path = "./examples/tests-master/references/well-typed" 
    folder_path = "../public-tests/week-3/main/secret/"
    # folder_path = "../public-tests/week-2/main/public/"
    binary_path = "./stella"  # Change this to the binary you want to execute
    binary_args = ["typecheck"]  # Add any arguments needed for the binary
    go_binary_path = "./my-stella"
    print("Started")
    process_files(folder_path, binary_path, binary_args, go_binary_path)
