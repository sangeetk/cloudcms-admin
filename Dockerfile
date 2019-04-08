FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD cloudcms-admin cloudcms-admin
ADD static static
ADD views views
EXPOSE 8080
CMD ["/cloudcms-admin"]
