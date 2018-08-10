#! /usr/bin/env bash
deb_ver=5.1.0-beta1
rpm_ver=5.1.0-beta1

wget https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform_${deb_ver}_amd64.deb

package_cloud push logdisplayplatform/testing/debian/jessie logdisplayplatform_${deb_ver}_amd64.deb
package_cloud push logdisplayplatform/testing/debian/wheezy logdisplayplatform_${deb_ver}_amd64.deb
package_cloud push logdisplayplatform/testing/debian/stretch logdisplayplatform_${deb_ver}_amd64.deb

wget https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform-${rpm_ver}.x86_64.rpm

package_cloud push logdisplayplatform/testing/el/6 logdisplayplatform-${rpm_ver}.x86_64.rpm
package_cloud push logdisplayplatform/testing/el/7 logdisplayplatform-${rpm_ver}.x86_64.rpm

rm logdisplayplatform*.{deb,rpm}
