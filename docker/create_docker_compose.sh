#!/bin/bash

blocks_dir=blocks
docker_dir=docker
template_dir=templates

logdisplayplatform_config_file=conf.tmp
logdisplayplatform_config=config

compose_header_file=compose_header.yml
fig_file=docker-compose.yaml
fig_config=docker-compose.yaml

if [ "$#" == 0 ]; then
    blocks=`ls $blocks_dir`
    if [ -z "$blocks" ]; then
        echo "No Blocks available in $blocks_dir"
    else
        echo "Available Blocks:"
        for block in $blocks; do
            echo "    $block"
        done
    fi
    exit 0
fi

for file in $logdisplayplatform_config_file $fig_file; do
    if [ -e $file ]; then
        echo "Deleting $file"
        rm $file
    fi
done

echo "Adding Compose header to $fig_file"
cat $compose_header_file >> $fig_file

for dir in $@; do
    current_dir=$blocks_dir/$dir
    if [ ! -d "$current_dir" ]; then
        echo "$current_dir is not a directory"
        exit 1
    fi

    if [ -e $current_dir/$logdisplayplatform_config ]; then
        echo "Adding $current_dir/$logdisplayplatform_config to $logdisplayplatform_config_file"
        cat $current_dir/$logdisplayplatform_config >> $logdisplayplatform_config_file
        echo "" >> $logdisplayplatform_config_file
    fi

    if [ -e $current_dir/$fig_config ]; then
        echo "Adding $current_dir/$fig_config to $fig_file"
        cat $current_dir/$fig_config >> $fig_file
        echo "" >> $fig_file
    fi
done
