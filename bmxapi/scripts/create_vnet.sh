#!/bin/bash

# Variables
RESOURCE_GROUP="myResourceGroup"
LOCATION="eastus"
VNET_NAME="myVNet"
SUBNET_NAME="mySubnet"

# Create Resource Group
az group create --name $RESOURCE_GROUP --location $LOCATION

# Create Virtual Network
az network vnet create --resource-group $RESOURCE_GROUP --name $VNET_NAME --address-prefix 10.0.0.0/16 --subnet-name $SUBNET_NAME --subnet-prefix 10.0.0.0/24

# Output VNet details
az network vnet show --resource-group $RESOURCE_GROUP --name $VNET_NAME
