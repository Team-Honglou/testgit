#! /usr/bin/env bash
version=5.0.2

wget https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform_${version}_amd64.deb

package_cloud push logdisplayplatform/stable/debian/jessie logdisplayplatform_${version}_amd64.deb
package_cloud push logdisplayplatform/stable/debian/wheezy logdisplayplatform_${version}_amd64.deb
package_cloud push logdisplayplatform/stable/debian/stretch logdisplayplatform_${version}_amd64.deb

package_cloud push logdisplayplatform/testing/debian/jessie logdisplayplatform_${version}_amd64.deb
package_cloud push logdisplayplatform/testing/debian/wheezy logdisplayplatform_${version}_amd64.deb --verbose
package_cloud push logdisplayplatform/testing/debian/stretch logdisplayplatform_${version}_amd64.deb --verbose

wget https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform-${version}-1.x86_64.rpm

package_cloud push logdisplayplatform/testing/el/6 logdisplayplatform-${version}-1.x86_64.rpm --verbose
package_cloud push logdisplayplatform/testing/el/7 logdisplayplatform-${version}-1.x86_64.rpm --verbose

package_cloud push logdisplayplatform/stable/el/7 logdisplayplatform-${version}-1.x86_64.rpm --verbose
package_cloud push logdisplayplatform/stable/el/6 logdisplayplatform-${version}-1.x86_64.rpm --verbose

rm logdisplayplatform*.{deb,rpm}
