package cgroups

import (
	"github.com/xianlubird/mydocker/cgroups/subsystems"
	log "github.com/Sirupsen/logrus"
)

type CgroupManager struct {
	Path string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager {
		Path: path,
	}
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range (subsystems.SubsystemIns) {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _,subSysIns := range (subsystems.SubsystemIns) {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

func (c *CgroupManager) Destory() error {
	for _, subSysIns := range (subsystems.SubsystemIns) {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}
