#!/bin/bash

# 1. 修复 ProcessAnnot.go 中的 int 类型 bug
python3 << 'PY1'
with open('annotation/ProcessAnnot.go', 'r') as f:
    content = f.read()

old = '''						} else if parm.ParmKind == reflect.Int {
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							getInt := c.Query(v.GenComment.Parms[index].ParmName)
							_ = value.Elem().Interface().(int)
							atoi, err := strconv.Atoi(getInt)
							if err != nil {
								fmt.Println(err)
							}
							_ = atoi
							parm.Value = value.Elem()
						} else if parm.ParmKind == reflect.String {'''

new = '''						} else if parm.ParmKind == reflect.Int {
							value := reflect.New(v.GenComment.Parms[index].ParmType)
							getInt := c.Query(v.GenComment.Parms[index].ParmName)
							atoi, err := strconv.Atoi(getInt)
							if err != nil {
								fmt.Println(err)
							}
							value.Elem().SetInt(int64(atoi))
							parm.Value = value.Elem()
						} else if parm.ParmKind == reflect.String {'''

content = content.replace(old, new)
with open('annotation/ProcessAnnot.go', 'w') as f:
    f.write(content)
print("1. ProcessAnnot.go fixed")
PY1

# 2. 修复 temroute.go 参数名和导入问题
python3 << 'PY2'
import re
with open('routers/temroute.go', 'r') as f:
    lines = f.readlines()

# 修复前3个参数名
count = 0
new_lines = []
for line in lines:
    if count < 3:
        if '"name"' in line and count == 0:
            line = line.replace('"name"', '"parm1"')
            count += 1
        elif '"password"' in line and count == 1:
            line = line.replace('"password"', '"parm2"')
            count += 1
        elif '"age"' in line and count == 2:
            line = line.replace('"age"', '"parm3"')
            count += 1
    new_lines.append(line)

content = ''.join(new_lines)

# 修复导入 - 删除重复的 utils 和空的 fmt
in_import = False
import_lines = []
other_lines = []
skip_utils = False

for line in content.split('\n'):
    if 'import (' in line:
        in_import = True
        import_lines.append(line)
        continue
    if in_import:
        if line.strip() == ')':
            in_import = False
            import_lines.append(line)
            continue
        # 跳过第二个 utils 声明
        if 'utils "' in line:
            if skip_utils:
                continue
            skip_utils = True
        # 跳过空的 fmt
        if line.strip() == '"fmt"':
            continue
        import_lines.append(line)
    else:
        other_lines.append(line)

result = '\n'.join(import_lines + other_lines)
with open('routers/temroute.go', 'w') as f:
    f.write(result)
print("2. temroute.go fixed")
PY2

echo "Done! Now building..."
