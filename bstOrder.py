# 二叉树遍历

class BinTreeNode(object):
    def __init__(self, data, left=None, right=None):
        self.data, self.left, self.right = data, left, right


class BinTree(object):
    def __init__(self, root=None):
        self.root = root

    @classmethod
    def build_from(cls, node_list):
        """通过节点信息构造二叉树
        第一次遍历我们构造 node 节点
        第二次遍历我们给 root 和 孩子赋值
        最后我们用 root 初始化这个类并返回一个对象
        """
        node_dict = {}
        for node_data in node_list:
            data = node_data['data']
            node_dict[data] = BinTreeNode(data)
        for node_data in node_list:
            data = node_data['data']
            node = node_dict[data]
            if node_data['is_root']:
                root = node
            node.left = node_dict.get(node_data['left'])
            node.right = node_dict.get(node_data['right'])
        return cls(root)

    def preorder(self, subtree):
        """
        前序遍历
        """
        if subtree is not None:
            print(subtree.data)
            self.preorder(subtree.left)
            self.preorder(subtree.right)

    def inorder(self, subtree):
        """
        中序遍历
        """
        if subtree is not None:
            self.inorder(subtree.left)
            print(subtree.data)
            self.inorder(subtree.right)

    def postorder(self, subtree):
        """
        后序遍历
        """
        if subtree is not None:
            self.postorder(subtree.left)
            self.postorder(subtree.right)
            print(subtree.data)

    def layer(self, subtree):
        """
        层序遍历
        """
        queue = [subtree]
        while queue:
            n = len(queue)
            for i in range(n):
                q = queue.pop(0)
                if q:
                    print(q.data)
                    queue.append(q.left if q.left else None)
                    queue.append(q.right if q.right else None)

    def maxDepth(self, subtree):
        """
        最大深度
        """
        if subtree is None:
            return 0
        return 1 + max(
            self.maxDepth(subtree.left), self.maxDepth(subtree.right))

    def minDepth(self, subtree):
        """
        最小深度
        """
        if subtree is None:
            return 0
        if subtree.left and subtree.right:
            return min(
                self.minDepth(subtree.left), self.minDepth(subtree.right)) + 1
        else:
            return max(
                self.minDepth(subtree.left), self.minDepth(subtree.right)) + 1

    def isbanlance(self, subtree):
        if subtree is None:
            return True
        left = self.maxDepth(subtree.left)
        right = self.maxDepth(subtree.right)
        return abs(left - right) <= 1 and self.isbanlance(
            subtree.left) and self.isbanlance(subtree.right)

    def isSymmetric(self, subtree):
        if not subtree.data:
            return True
        return self.isMirror(subtree.left, subtree.right)

    def isMirror(self, p, q):
        if not p and not q:
            return True
        if not p or not q:
            return False
        le = self.isMirror(p.left, q.right)
        r = self.isMirror(p.right, q.left)
        return p.data == q.data and le and r


node_list = [
    {'data': 'A', 'left': 'B', 'right': 'C', 'is_root': True},
    {'data': 'B', 'left': 'D', 'right': 'E', 'is_root': False},
    {'data': 'D', 'left': None, 'right': None, 'is_root': False},
    {'data': 'E', 'left': 'H', 'right': None, 'is_root': False},
    {'data': 'H', 'left': None, 'right': None, 'is_root': False},
    {'data': 'C', 'left': 'F', 'right': 'G', 'is_root': False},
    {'data': 'F', 'left': None, 'right': None, 'is_root': False},
    {'data': 'G', 'left': 'I', 'right': 'J', 'is_root': False},
    {'data': 'I', 'left': None, 'right': None, 'is_root': False},
    {'data': 'J', 'left': None, 'right': None, 'is_root': False},


]

btree = BinTree.build_from(node_list)
# btree.preorder(btree.root)
# btree.inorder(btree.root)
# btree.postorder(btree.root)
# btree.layer(btree.root)
# print(btree.maxDepth(btree.root))
# print(btree.minDepth(btree.root))
# print(btree.isbanlance(btree.root))
print(btree.isSymmetric(btree.root))
