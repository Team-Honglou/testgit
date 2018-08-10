  #!/bin/bash

_circle_token=$1
_logdisplayplatform_version=$2

trigger_build_url=https://circleci.com/api/v1/project/logdisplayplatform/logdisplayplatform-docker/tree/master?circle-token=${_circle_token}

post_data=$(cat <<EOF
{
  "build_parameters": {
    "GRAFANA_VERSION": "${_logdisplayplatform_version}"
  }
}
EOF
)

echo ${post_data}

curl \
--header "Accept: application/json" \
--header "Content-Type: application/json" \
--data "${post_data}" \
--request POST ${trigger_build_url}
