#!/usr/bin/zsh

# Get the current version
v=$(go list -mod=mod -m all | grep github.com/tektoncd/pipeline | awk '{print $2}')
echo ${v}

# Get the list of all versions
list=( $(go list -mod=mod -m -versions goa.design/goa/v3) )
echo ${list[-1]}

# Compare the versions
if [ "${v}" != "${list[-1]}" ]
then
  # goa.design/goa/v3 v3.7.12
  sed -i "s@goa.design/goa/v3 ${v}/goa.design/goa/v3 ${list[-1]}@g" go.mod
fi