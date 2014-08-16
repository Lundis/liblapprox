package expr

func (self *Node) OptimizeCommon() {
	self.Replace(common_constants)
	self.Optimize()
}

func (self *Node) Optimize() {

}