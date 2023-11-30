FROM projects.registry.vmware.com/vmware-cloud-director/cloud-director-named-disk-csi-driver:1.4.1

WORKDIR /opt/vcloud/bin

ADD bin/cloud-director-named-disk-csi-driver .
RUN chmod +x /opt/vcloud/bin/cloud-director-named-disk-csi-driver

# USER nobody
ENTRYPOINT ["/bin/bash", "-l", "-c"]
