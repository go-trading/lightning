gcloud compute ssh datasaver2 --command="sudo mkdir /datasaver"
gcloud compute scp  ./history/datasaver.service datasaver2:~
gcloud compute ssh datasaver2 --command="sudo mv ~/datasaver.service /etc/systemd/system/datasaver.service"
gcloud compute ssh datasaver2 --command="sudo chmod 644 /etc/systemd/system/datasaver.service"
gcloud compute ssh datasaver2 --command="sudo systemctl daemon-reload && sudo systemctl enable datasaver"
gcloud compute scp ./history/datasaver datasaver2:~
gcloud compute ssh datasaver2 --command="sudo mv ~/datasaver /datasaver/datasaver.new"
gcloud compute ssh datasaver2 --command="sudo systemctl start datasaver"



sudo cp ./history/datasaver.service /etc/systemd/system/datasaver.service
sudo chmod 644 /etc/systemd/system/datasaver.service
sudo systemctl daemon-reload && sudo systemctl enable datasaver
sudo systemctl start datasaver
systemctl status datasaver

sudo cp ./history/datasaver.service /etc/systemd/system/datasaver.service
sudo systemctl daemon-reload
sudo systemctl restart datasaver
systemctl status datasaver


gcloud beta compute --project=lightning-272813 instances create datasaver1 --zone=us-central1-a --machine-type=f1-micro --subnet=default --network-tier=PREMIUM --maintenance-policy=MIGRATE --service-account=6061868947-compute@developer.gserviceaccount.com --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append --tags=http-server --image=debian-9-stretch-v20200309 --image-project=debian-cloud --boot-disk-size=100GB --boot-disk-type=pd-standard --boot-disk-device-name=datasaver1 --reservation-affinity=any
gcloud beta compute --project=lightning-272813 instances create datasaver2 --zone=us-central1-a --machine-type=f1-micro --subnet=default --network-tier=PREMIUM --maintenance-policy=MIGRATE --service-account=6061868947-compute@developer.gserviceaccount.com --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append --tags=http-server --image=debian-9-stretch-v20200309 --image-project=debian-cloud --boot-disk-size=100GB --boot-disk-type=pd-standard --boot-disk-device-name=datasaver2 --reservation-affinity=any

gcloud compute --project=lightning-272813 firewall-rules create default-allow-http --direction=INGRESS --priority=1000 --network=default --action=ALLOW --rules=tcp:80 --source-ranges=0.0.0.0/0 --target-tags=http-server

