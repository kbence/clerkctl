package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"users"},
	Short:   "User management",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var userListCmdParams struct {
	limit int64
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing users",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancelFunc := context.WithTimeout(cmd.Context(), 5*time.Second)
		defer cancelFunc()

		usersLeft := userListCmdParams.limit
		offset := int64(0)

		for {
			limit := int64(Min(10, usersLeft))

			userList, err := user.List(ctx, &user.ListParams{
				ListParams: clerk.ListParams{
					Limit: &limit, Offset: &offset,
				},
			})
			if err != nil {
				return err
			}

			if len(userList.Users) == 0 {
				break
			}

			for _, user := range userList.Users {
				fmt.Println(user.ID, user.EmailAddresses[0].EmailAddress)
			}

			offset += int64(len(userList.Users))
			usersLeft -= int64(len(userList.Users))

			if usersLeft <= 0 {
				break
			}
		}

		return nil
	},
}

var userDeleteCmdParams struct {
	Email string
}

var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user(s) by ID or email address",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancelFunc := context.WithTimeout(cmd.Context(), 5*time.Second)
		defer cancelFunc()

		for _, userId := range args {
			if !strings.HasPrefix(userId, "user_") {
				userList, err := user.List(ctx, &user.ListParams{EmailAddresses: []string{userId}})

				if err != nil {
					return err
				}

				if len(userList.Users) != 1 {
					log.Printf("couldn't find an exact match to '%s'", userId)
					continue
				}

				userId = userList.Users[0].ID
			}

			deletedUser, err := user.Delete(ctx, userId)
			if err != nil {
				return err
			}

			log.Printf("deleted user %s", deletedUser.ID)
		}

		return nil
	},
}

func init() {
	userListCmd.Flags().Int64VarP(&userListCmdParams.limit, "limit", "m", 0, "Maximum number of users to list")
	userDeleteCmd.Flags().StringVarP(&userDeleteCmdParams.Email, "email", "e", "", "Delete user by email address")

	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userDeleteCmd)
	rootCmd.AddCommand(userCmd)
}
