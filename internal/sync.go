package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	git "gopkg.in/src-d/go-git.v4"
	gitconfig "gopkg.in/src-d/go-git.v4/config"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

type Sync struct {
	syncService SyncService
}

func NewSync() *Sync {
	viper.GetBool("init")
	return &Sync{}
}

func (s *Sync) PullChanges() {

}

func (s *Sync) Push() {
	r, err := git.PlainOpen(viper.GetString("library"))
	if err == git.ErrRepositoryNotExists {
		fmt.Println(err)
		r = initializeOrCloneRepository()
	}
	err = r.Fetch(&git.FetchOptions{
		Auth: getAuth(),
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		panic(fmt.Errorf("Fatal error fetching repository: %s \n", err))
	}

	if err == git.NoErrAlreadyUpToDate {
		fmt.Println(err)
	}
	err = r.CreateBranch(&gitconfig.Branch{
		Name: "testing2",
	})
	if err != nil && err != git.ErrBranchExists {
		panic(fmt.Errorf("Fatal error creating branch: %s \n", err))
	}
	if err == git.ErrBranchExists {
		fmt.Println(err)
	}
}

func initializeOrCloneRepository() *git.Repository {
	if len(viper.GetString("repository.remote")) > 0 {
		r, err := git.PlainClone(viper.GetString("library"), false, &git.CloneOptions{
			URL:  viper.GetString("repository.remote"),
			Auth: getAuth(),
		})
		if err != nil {
			panic(fmt.Errorf("Fatal error cloning repository: %s \n", err))
		}
		return r
	}
	r, err := git.PlainInit(viper.GetString("library"), false)
	if err != nil {
		panic(fmt.Errorf("Fatal error initializing repository: %s \n", err))
	}
	return r
}

func getAuth() *gitssh.PublicKeys {
	s := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
	sshKey, _ := ioutil.ReadFile(s)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth := &gitssh.PublicKeys{User: "git", Signer: signer}
	return auth
}
